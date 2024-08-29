package service

import (
	"context"
	"encoding/json"
	"log/slog"
	pb "sale/genproto/sale"
	"sale/storage"
)

type ProductKafkaService struct {
	logger *slog.Logger
	repo   storage.Istorage
}

func NewProductKafkaService(logger *slog.Logger, repo storage.Istorage) *ProductKafkaService {
	return &ProductKafkaService{
		logger: logger,
		repo:   repo,
	}
}

func (s ProductKafkaService) UpdateProduct(message []byte) {
	s.logger.Info("UpdateProduct rpc method is started")
	req := pb.UpdateProductRequest{}
	err := json.Unmarshal(message, &req)
	if err != nil {
		s.logger.Error(err.Error())
		return
	}

	err = s.repo.Product().UpdateProduct(context.Background(), &req)
	if err != nil {
		s.logger.Error(err.Error())
		return
	}
	s.logger.Info("UpdateProduct finished succesfully")
}

func (s ProductKafkaService) DeleteProduct(message []byte) {
	s.logger.Info("DeleteProduct rpc method is started")
	req := pb.ProductId{}
	err := json.Unmarshal(message, &req)
	if err != nil {
		s.logger.Error(err.Error())
		return
	}
	err = s.repo.Product().DeleteProduct(context.Background(), &req)
	if err != nil {
		s.logger.Error(err.Error())
		return
	}
	s.logger.Info("DeleteProduct finished succesfully")
}
