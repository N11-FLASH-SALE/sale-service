package postgres

import (
	"database/sql"
	"log/slog"
	logger "sale/logs"
	pb "sale/genproto/sale"

)

type FeedbackRepository struct {
	Db *sql.DB
	lg *slog.Logger
}

func NewFeedbackRepository(db *sql.DB) *FeedbackRepository {
	return &FeedbackRepository{Db: db, lg: logger.NewLogger()}
}

func (repo *FeedbackRepository) CreateFeedback(request *pb.CreateFeedbackRequest) (*pb.FeedbackResponse, error) {
	return nil, nil
}

func (repo *FeedbackRepository) GetFeedback(request *pb.GetFeedbackRequest) (*pb.GetFeedbackResponse, error) {
	return nil, nil
}

func (repo *FeedbackRepository) GetFeedbackOfUser(request *pb.GetFeedbackOfUserRequest) (*pb.GetFeedbackOfUserResponse, error) {
	return nil, nil
}
