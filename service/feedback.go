package service

import (
	"context"
	"log/slog"
	pb "sale/genproto/sale"
	"sale/storage"
)

type FeedbackService struct {
	pb.UnimplementedFeedbackServer
	logger *slog.Logger
	repo   storage.Istorage
}

func NewFeedbackService(logger *slog.Logger, repo storage.Istorage) *FeedbackService {
	return &FeedbackService{
		logger: logger,
		repo:   repo,
	}
}

func (s FeedbackService) CreateFeedback(ctx context.Context, req *pb.CreateFeedbackRequest) (*pb.FeedbackResponse, error) {
	s.logger.Info("CreateFeedback rpc method is started")
	res, err := s.repo.Feedback().CreateFeedback(ctx, req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	s.logger.Info("CreateFeedback finished succesfully")
	return res, nil
}

func (s FeedbackService) GetFeedbackOfUser(ctx context.Context, req *pb.GetFeedbackOfUserRequest) (*pb.GetFeedbackOfUserResponse, error) {
	s.logger.Info("GetFeedbackOfUser rpc method is started")
	res, err := s.repo.Feedback().GetFeedbackOfUser(ctx, req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	s.logger.Info("GetFeedbackOfUser finished succesfully")
	return res, nil
}

func (s FeedbackService) GetFeedback(ctx context.Context, request *pb.GetFeedbackRequest) (*pb.GetFeedbackResponse, error) {
	s.logger.Info("GetFeedback rpc method is started")
	res, err := s.repo.Feedback().GetFeedback(ctx, request)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	s.logger.Info("GetFeedback finished succesfully")
	return res, nil
}
