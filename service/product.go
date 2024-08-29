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

func (s ProductService) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.Void, error) {
	s.logger.Info("UpdateProduct rpc method is started")
	err := s.repo.Product().UpdateProduct(ctx, req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	s.logger.Info("UpdateProduct finished succesfully")
	return &pb.Void{}, nil
}

func (s ProductService) DeleteProduct(ctx context.Context, req *pb.ProductId) (*pb.Void, error) {
	s.logger.Info("DeleteProduct rpc method is started")
	err := s.repo.Product().DeleteProduct(ctx, req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	s.logger.Info("DeleteProduct finished succesfully")
	return &pb.Void{}, nil
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
