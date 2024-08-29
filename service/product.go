package service

import (
	"log/slog"
	pb "sale/genproto/sale"
	"sale/storage"
)

type HealthMonitoringService struct {
	pb.UnimplementedProductServer
	logger *slog.Logger
	repo   storage.Istorage
}

func NewHealthMonitoringService(logger *slog.Logger, repo storage.Istorage) *HealthMonitoringService {
	return &HealthMonitoringService{
		logger: logger,
		repo:   repo,
	}
}
