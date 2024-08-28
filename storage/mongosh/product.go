package mongos

import (
	"context"
	pb "sale/genproto/sale"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductsRepo struct {
	Coll *mongo.Collection
}

func NewProductsRepository(db *mongo.Database) *ProductsRepo {
	return &ProductsRepo{Coll: db.Collection("product")}
}

func (r *ProductsRepo) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.ProductId, error) {

	price := req.Price - req.Price*(float64(req.Stock)/100)

	product := bson.D{
		{Key: "name", Value: req.Name},
		{Key: "description", Value: req.Description},
		{Key: "price", Value: req.Price},
		{Key: "stock", Value: req.Stock},
		{Key: "price_with_stock", Value: price},
		{Key: "limit_of_product", Value: req.LimitOfProduct},
		{Key: "size", Value: req.Size},
		{Key: "color", Value: req.Color},
		{Key: "start_date", Value: req.StartDate},
		{Key: "end_date", Value: req.EndDate},
		{Key: "seller_id", Value: req.SellerId},
	}

	result, err := r.Coll.InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	return &pb.ProductId{Id: id}, nil
}
