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
	priceWithStock := req.Price - req.Price*stockPercentage

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
		{Key: "price", Value: req.Price},
		{Key: "stock", Value: req.Stock},
		{Key: "price_with_stock", Value: priceWithStock},
		{Key: "limit_of_product", Value: req.LimitOfProduct},
		{Key: "size", Value: req.Size},
		{Key: "color", Value: req.Color},
		{Key: "start_date", Value: startDate},
		{Key: "end_date", Value: endDate},
		{Key: "seller_id", Value: req.SellerId},
	}

	result, err := r.Coll.InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	return &pb.ProductId{Id: id}, nil
}

func (r *ProductsRepo) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	filter := bson.M{}

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
	if req.LimitOfProduct > 0 {
		filter["limit_of_product"] = bson.M{"$gte": req.LimitOfProduct}
	}
	if req.PriceWithStock > 0 {
		filter["price_with_stock"] = bson.M{"$gte": req.PriceWithStock}
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
			Id:             product.ID.Hex(),
			Name:           product.Name,
			Description:    product.Description,
			Price:          product.Price,
			Stock:          product.Stock,
			PriceWithStock: product.PriceWithStock,
			LimitOfProduct: product.LimitOfProduct,
			Size:           product.Size,
			Color:          product.Color,
			StartDate:      product.StartDate.Format("2006-01-02"),
			EndDate:        product.EndDate.Format("2006-01-02"),
			SellerId:       product.SellerID,
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
