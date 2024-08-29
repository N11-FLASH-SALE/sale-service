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
