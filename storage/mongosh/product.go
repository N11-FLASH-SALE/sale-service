package mongos

import (
	"context"
	"fmt"
	pb "sale/genproto/sale"
	"sale/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductsRepo struct {
	Coll *mongo.Collection
}

func NewProductsRepository(db *mongo.Database) *ProductsRepo {
	return &ProductsRepo{Coll: db.Collection("product")}
}

func (r *ProductsRepo) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.ProductId, error) {

	stockPercentage := float64(req.Stock) / 100
	priceWithStock := req.PriceWithoutStock - req.PriceWithoutStock*stockPercentage

	startDate, err := time.Parse("2006-01-02", req.GetStartDate())
	if err != nil {
		return nil, fmt.Errorf("invalid start_date format: %v", err)
	}
	endDate, err := time.Parse("2006-01-02", req.GetEndDate())
	if err != nil {
		return nil, fmt.Errorf("invalid end_date format: %v", err)
	}

	product := bson.D{
		{Key: "name", Value: req.Name},
		{Key: "description", Value: req.Description},
		{Key: "price", Value: priceWithStock},
		{Key: "stock", Value: req.Stock},
		{Key: "price_without_stock", Value: req.PriceWithoutStock},
		{Key: "limit_of_product", Value: req.LimitOfProduct},
		{Key: "size", Value: req.Size},
		{Key: "color", Value: req.Color},
		{Key: "start_date", Value: startDate},
		{Key: "end_date", Value: endDate},
		{Key: "seller_id", Value: req.SellerId},
		{Key: "photos", Value: []string{}},
		{Key: "deleted_at", Value: nil},
	}

	result, err := r.Coll.InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	return &pb.ProductId{Id: id}, nil
}

func (r *ProductsRepo) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	filter := bson.M{
		"deleted_at": nil,
	}

	if req.Name != "" {
		filter["name"] = bson.M{"$regex": req.Name, "$options": "i"}
	}
	if req.MinPrice > 0 {
		filter["price"] = bson.M{"$gte": req.MinPrice}
	}
	if req.MaxPrice > 0 {
		if val, ok := filter["price"]; ok {
			val.(bson.M)["$lte"] = req.MaxPrice
		} else {
			filter["price"] = bson.M{"$lte": req.MaxPrice}
		}
	}
	if req.Stock > 0 {
		filter["stock"] = bson.M{"$gte": req.Stock}
	}
	if req.SellerId != "" {
		filter["seller_id"] = req.SellerId
	}

	// Filter out products with an expired end_date
	now := time.Now()
	filter["end_date"] = bson.M{"$gte": now}

	// Set limit and offset for pagination
	options := options.Find()
	if req.Limit > 0 {
		options.SetLimit(req.Limit)
	}
	if req.Offset > 0 {
		options.SetSkip(req.Offset)
	}

	// Query the MongoDB collection
	cursor, err := r.Coll.Find(ctx, filter, options)
	if err != nil {
		return nil, fmt.Errorf("failed to find products: %w", err)
	}
	defer cursor.Close(ctx)

	// Prepare the response
	var products []models.Product
	if err = cursor.All(ctx, &products); err != nil {
		return nil, fmt.Errorf("failed to decode products: %w", err)
	}

	// Map the products to the protobuf message
	var pbProducts []*pb.Products
	for _, product := range products {
		pbProducts = append(pbProducts, &pb.Products{
			Id:                product.ID.Hex(),
			Name:              product.Name,
			Description:       product.Description,
			Price:             product.Price,
			Stock:             product.Stock,
			PriceWithoutStock: product.PriceWithOutStock,
			LimitOfProduct:    product.LimitOfProduct,
			Size:              product.Size,
			Color:             product.Color,
			StartDate:         product.StartDate.Format("2006-01-02"),
			EndDate:           product.EndDate.Format("2006-01-02"),
			SellerId:          product.SellerID,
			Photos:            product.Photos,
		})
	}

	// Count total number of matching products
	totalCount, err := r.Coll.CountDocuments(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to count products: %w", err)
	}

	return &pb.GetProductResponse{
		Product:    pbProducts,
		TotalCount: totalCount,
	}, nil
}

func (r *ProductsRepo) GetProductById(ctx context.Context, req *pb.ProductId) (*pb.GetProductByIdResponse, error) {
	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid product id: %w", err)
	}

	filter := bson.D{
		{Key: "_id", Value: objID},
		{Key: "deleted_at", Value: nil},
	}

	var product models.Product
	err = r.Coll.FindOne(ctx, filter).Decode(&product)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("product not found")
	} else if err != nil {
		return nil, fmt.Errorf("failed to find product: %w", err)
	}
	return &pb.GetProductByIdResponse{
		Id:                product.ID.Hex(),
		Name:              product.Name,
		Description:       product.Description,
		Price:             product.Price,
		Stock:             product.Stock,
		PriceWithoutStock: product.PriceWithOutStock,
		LimitOfProduct:    product.LimitOfProduct,
		Size:              product.Size,
		Color:             product.Color,
		StartDate:         product.StartDate.Format("2006-01-02"),
		EndDate:           product.EndDate.Format("2006-01-02"),
		SellerId:          product.SellerID,
		Photos:            product.Photos,
	}, nil
}

func (r *ProductsRepo) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) error {
	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return fmt.Errorf("invalid product id: %w", err)
	}
	update := bson.M{}

	if req.Name != "" {
		update["name"] = req.Name
	}
	if req.Description != "" {
		update["description"] = req.Description
	}
	if req.Price > 0 {
		update["price"] = req.Price
	}
	if req.Stock > 0 {
		update["stock"] = req.Stock
	}
	if req.PriceWithoutStock > 0 {
		update["price_without_stock"] = req.PriceWithoutStock
	}
	if req.LimitOfProduct > 0 {
		update["limit_of_product"] = req.LimitOfProduct
	}
	if len(req.Size) > 0 {
		update["size"] = req.Size
	}
	if len(req.Color) > 0 {
		update["color"] = req.Color
	}
	if req.StartDate != "" {
		startDate, err := time.Parse("2006-01-02", req.StartDate)
		if err != nil {
			return fmt.Errorf("invalid start_date format: %v", err)
		}
		update["start_date"] = startDate
	}
	if req.EndDate != "" {
		endDate, err := time.Parse("2006-01-02", req.EndDate)
		if err != nil {
			return fmt.Errorf("invalid end_date format: %v", err)
		}
		update["end_date"] = endDate
	}

	filter := bson.M{"_id": objID, "deleted_at": nil}
	updateResult, err := r.Coll.UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}

	if updateResult.MatchedCount == 0 {
		return fmt.Errorf("product not found or already deleted")
	}

	return nil
}

func (r *ProductsRepo) DeleteProduct(ctx context.Context, req *pb.ProductId) error {
	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return fmt.Errorf("invalid product id: %w", err)
	}

	filter := bson.M{"_id": objID, "deleted_at": nil}
	update := bson.M{"$set": bson.M{"deleted_at": time.Now()}}
	updateResult, err := r.Coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	if updateResult.MatchedCount == 0 {
		return fmt.Errorf("product not found or already deleted")
	}

	return nil
}

func (r *ProductsRepo) IsProductOk(ctx context.Context, req *pb.ProductId) error {
	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return fmt.Errorf("invalid product id: %w", err)
	}
	now := time.Now()
	filter := bson.M{
		"_id":        objID,
		"end_date":   bson.M{"$gte": now},
		"deleted_at": nil,
	}

	var product models.Product
	err = r.Coll.FindOne(ctx, filter).Decode(&product)
	if err == mongo.ErrNoDocuments {
		return fmt.Errorf("product not found or expired or deleted")
	} else if err != nil {
		return fmt.Errorf("failed to check product status: %w", err)
	}

	return nil
}

func (r *ProductsRepo) AddPhotosToProduct(ctx context.Context, req *pb.AddPhotosRequest) error {
	objID, err := primitive.ObjectIDFromHex(req.ProductId)
	if err != nil {
		return fmt.Errorf("invalid product id: %w", err)
	}

	update := bson.M{
		"$push": bson.M{"photos": req.PhotoUrl},
	}

	filter := bson.M{
		"_id":        objID,
		"deleted_at": nil,
	}

	result, err := r.Coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("product not found or has been deleted")
	}

	return nil
}

func (r *ProductsRepo) DeletePhotosFromProduct(ctx context.Context, req *pb.DeletePhotosRequest) error {
	objID, err := primitive.ObjectIDFromHex(req.ProductId)
	if err != nil {
		return fmt.Errorf("invalid product id: %w", err)
	}

	update := bson.M{
		"$pull": bson.M{"photos": req.PhotoUrl},
	}

	filter := bson.M{
		"_id":        objID,
		"deleted_at": nil,
	}

	result, err := r.Coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("product not found or has been deleted")
	}

	return nil
}
