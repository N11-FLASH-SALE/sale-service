package service

import (
	"context"
	"log/slog"
	pb "sale/genproto/sale"
	"sale/storage"
)

type BoughtService struct {
	pb.UnimplementedBoughtServer
	logger *slog.Logger
	repo   storage.Istorage
}

func NewBoughtService(logger *slog.Logger, repo storage.Istorage) *BoughtService {
	return &BoughtService{
		logger: logger,
		repo:   repo,
	}
}

func (s BoughtService) CreateBought(ctx context.Context, req *pb.CreateBoughtRequest) (*pb.BoughtResponse, error) {
	s.logger.Info("CreateBought rpc method is started")
	res, err := s.repo.Bought().CreateBought(ctx, req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	s.logger.Info("CreateBought finished succesfully")
	return res, nil
}

func (s BoughtService) GetBought(ctx context.Context, req *pb.GetBoughtRequest) (*pb.GetBoughtResponse, error) {
	s.logger.Info("GetBought rpc method is started")
	res, err := s.repo.Bought().GetBought(ctx, req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	s.logger.Info("GetBought finished succesfully")
	return res, nil
}

func (s BoughtService) GetBoughtOfUser(ctx context.Context, req *pb.GetBoughtOfUserRequest) (*pb.GetBoughtOfUserResponse, error) {
	s.logger.Info("GetBoughtOfUser rpc method is started")
	res, err := s.repo.Bought().GetBoughtOfUser(ctx, req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	s.logger.Info("GetBoughtOfUser finished succesfully")
	return res, nil
}
