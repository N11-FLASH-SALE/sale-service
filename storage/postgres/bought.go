package postgres

import (
	"context"
	"database/sql"
	"log/slog"
	logger "sale/logs"
	pb "sale/genproto/sale"
)

type BougthRepository struct {
	Db *sql.DB
	lg *slog.Logger
}

func NewBougthRepository(db *sql.DB) *BougthRepository {
	return &BougthRepository{Db: db, lg: logger.NewLogger()}
}

func (repo *BougthRepository) CreateBought(ctx context.Context, req *pb.CreateBoughtRequest) (*pb.BoughtResponse, error){
	return nil, nil
}

func (repo *BougthRepository) GetBought(ctx context.Context, req *pb.GetBoughtRequest) (*pb.GetBoughtResponse, error){
	return nil, nil
}

func (repo *BougthRepository) GetBoughtOfUser(ctx context.Context, req *pb.GetBoughtOfUserRequest) (*pb.GetBoughtOfUserResponse, error){
	return nil, nil
}