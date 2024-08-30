package service

import (
	"context"
	"log/slog"
	pb "sale/genproto/sale"
	"sale/storage"
)

type ProcessService struct {
	pb.UnimplementedProcessServer
	logger *slog.Logger
	repo   storage.Istorage
}

func NewProcessService(logger *slog.Logger, repo storage.Istorage) *ProcessService {
	return &ProcessService{
		logger: logger,
		repo:   repo,
	}
}

func (s ProcessService) CreateProcess(ctx context.Context, req *pb.CreateProcessRequest) (*pb.ProcessResponse, error) {
	s.logger.Info("CreateProcess rpc method is started")
	res, err := s.repo.Processes().CreateProcess(ctx, req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	s.logger.Info("CreateProcess finished succesfully")
	return res, nil
}

func (s ProcessService) GetProcessOfUserByProductId(ctx context.Context, req *pb.GetProcessOfUserByProductIdRequest) (*pb.GetProcessOfUserByProductIdResponse, error) {
	s.logger.Info("GetProcessOfUser rpc method is started")
	res, err := s.repo.Processes().GetProcessOfUserByProductId(ctx, req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	s.logger.Info("GetProcessOfUser finished succesfully")
	return res, nil
}

func (s ProcessService) GetProcessByProductId(ctx context.Context, req *pb.GetProcessByProductIdRequest) (*pb.GetProcessByProductIdResponse, error) {
	s.logger.Info("GetProcessByProductId rpc method is started")
	res, err := s.repo.Processes().GetProcessByProductId(ctx, req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	s.logger.Info("GetProcessByProductId finished succesfully")
	return res, nil
}

func (s ProcessService) UpdateProcess(ctx context.Context, req *pb.UpdateProcessRequest) (*pb.Void, error) {
	s.logger.Info("GetProcessByUserId rpc method is started")
	err := s.repo.Processes().UpdateProcess(ctx, req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	s.logger.Info("GetProcessByUserId finished succesfully")
	return &pb.Void{}, nil
}

func (s ProcessService) CancelProcess(ctx context.Context, req *pb.CancelProcessRequest) (*pb.CancelProcessResponse, error) {
	s.logger.Info("CancelProcess rpc method is started")
	res, err := s.repo.Processes().CancelProcess(ctx, req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	s.logger.Info("CancelProcess finished succesfully")
	return res, nil
}

func (s ProcessService) GetProcessByUserId(ctx context.Context, req *pb.GetProcessByUserIdRequest) (*pb.GetProcessByUserIdResponse, error) {
	s.logger.Info("GetProcessByUserId rpc method is started")
	res, err := s.repo.Processes().GetProcessByUserId(ctx, req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	s.logger.Info("GetProcessByUserId finished succesfully")
	return res, nil
}

func (s ProcessService) GetProcessById(ctx context.Context, req *pb.GetProcessByIdRequest) (*pb.GetProcessByIdResponse, error) {
	s.logger.Info("GetProcessById rpc method is started")
	res, err := s.repo.Processes().GetProcessById(ctx, req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	s.logger.Info("GetProcessById finished succesfully")
	return res, nil
}
