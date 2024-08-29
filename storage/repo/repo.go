package repo

import (
	"context"
	pb "sale/genproto/sale"
)

type Product interface {
	CreateProduct(context.Context, *pb.CreateProductRequest) (*pb.ProductId, error)
	GetProduct(context.Context, *pb.GetProductRequest) (*pb.GetProductResponse, error)
	GetProductById(context.Context, *pb.ProductId) (*pb.GetProductByIdResponse, error)
	UpdateProduct(context.Context, *pb.UpdateProductRequest) error
	DeleteProduct(context.Context, *pb.ProductId) error
	IsProductOk(context.Context, *pb.ProductId) error
	AddPhotosToProduct(context.Context, *pb.AddPhotosRequest) error
	DeletePhotosFromProduct(context.Context, *pb.DeletePhotosRequest) error
}

type Bought interface {
	CreateBought(context.Context, *pb.CreateBoughtRequest) (*pb.BoughtResponse, error)
	GetBought(context.Context, *pb.GetBoughtRequest) (*pb.GetBoughtResponse, error)
	GetBoughtOfUser(context.Context, *pb.GetBoughtOfUserRequest) (*pb.GetBoughtOfUserResponse, error)
}

type Feedback interface {
	CreateFeedback(context.Context, *pb.CreateFeedbackRequest) (*pb.FeedbackResponse, error)
	GetFeedback(context.Context, *pb.GetFeedbackRequest) (*pb.GetFeedbackResponse, error)
	GetFeedbackOfUser(context.Context, *pb.GetFeedbackOfUserRequest) (*pb.GetFeedbackOfUserResponse, error)
}

type Processes interface {
	CreateProcess(context.Context, *pb.CreateProcessRequest) (*pb.ProcessResponse, error)
	GetProcessOfUserByProductId(context.Context, *pb.GetProcessOfUserByProductIdRequest) (*pb.GetProcessOfUserByProductIdResponse, error)
	GetProcessByProductId(context.Context, *pb.GetProcessByProductIdRequest) (*pb.GetProcessByProductIdResponse, error)
	UpdateProcess(context.Context, *pb.UpdateProcessRequest) error
	CancelProcess(context.Context, *pb.CancelProcessRequest) error
}

type Wishlist interface {
	CreateWishlist(context.Context, *pb.CreateWishlistRequest) (*pb.WishlistResponse, error)
	GetWishlist(context.Context, *pb.GetWishlistRequest) (*pb.GetWishlistResponse, error)
	GetWishlistById(context.Context, *pb.GetWishlistByIdRequest) (*pb.GetWishlistByIdResponse, error)
}
