package service

import (
	"context"
	"log/slog"
	pb "sale/genproto/sale"
	"sale/storage"
)

type ProductService struct {
	pb.UnimplementedProductServer
	logger *slog.Logger
	repo   storage.Istorage
}

func NewProductService(logger *slog.Logger, repo storage.Istorage) *ProductService {
	return &ProductService{
		logger: logger,
		repo:   repo,
	}
}

func (s ProductService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.ProductId, error) {
	s.logger.Info("CreateProduct rpc method is started")
	res, err := s.repo.Product().CreateProduct(ctx, req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	s.logger.Info("CreateProduct finished succesfully")
	return res, nil
}

func (s ProductService) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	s.logger.Info("GetProduct rpc method is started")
	res, err := s.repo.Product().GetProduct(ctx, req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	s.logger.Info("GetProduct finished succesfully")
	return res, nil
}

func (s ProductService) GetProductById(ctx context.Context, req *pb.ProductId) (*pb.GetProductByIdResponse, error) {
	s.logger.Info("GetProductById rpc method is started")
	res, err := s.repo.Product().GetProductById(ctx, req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	s.logger.Info("GetProductById finished succesfully")
	return res, nil
}

func (s ProductService) GetProductsByUserId(ctx context.Context, req *pb.GetProductsByUserIdRequest) (*pb.GetProductsByUserIdResponse, error) {
	s.logger.Info("GetProductsByUserId rpc method is started")
	res, err := s.repo.Product().GetProductsByUserId(ctx, req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	s.logger.Info("GetProductsByUserId finished succesfully")
	return res, nil
}

func (s ProductService) IsProductOk(ctx context.Context, req *pb.ProductId) (*pb.Void, error) {
	s.logger.Info("IsProductOk rpc method is started")
	err := s.repo.Product().IsProductOk(ctx, req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	s.logger.Info("IsProductOk finished succesfully")
	return &pb.Void{}, nil
}

func (s ProductService) AddPhotosToProduct(ctx context.Context, req *pb.AddPhotosRequest) (*pb.Void, error) {
	s.logger.Info("AddPhotosToProduct rpc method is started")
	err := s.repo.Product().AddPhotosToProduct(ctx, req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	s.logger.Info("AddPhotosToProduct finished succesfully")
	return &pb.Void{}, nil
}

func (s ProductService) DeletePhotosFromProduct(ctx context.Context, req *pb.DeletePhotosRequest) (*pb.Void, error) {
	s.logger.Info("DeletePhotosFromProduct rpc method is started")
	err := s.repo.Product().DeletePhotosFromProduct(ctx, req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	s.logger.Info("DeletePhotosFromProduct finished succesfully")
	return &pb.Void{}, nil
}

func (s ProductService) UpdateLimitOfProduct(ctx context.Context, req *pb.UpdateLimitOfProductRequest) (*pb.Void, error) {
	s.logger.Info("UpdateLimitOfProduct rpc method is started")
	err := s.repo.Product().UpdateLimitOfProduct(ctx, req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	s.logger.Info("UpdateLimitOfProduct finished succesfully")
	return &pb.Void{}, nil
}

func (s ProductService) IsProductExists(ctx context.Context, req *pb.ProductId) (*pb.Void, error) {
	s.logger.Info("IsProductExists rpc method is started")
	err := s.repo.Product().IsProductExists(ctx, req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	s.logger.Info("IsProductExists finished succesfully")
	return &pb.Void{}, nil
}
