package service

import (
	"context"
	"log/slog"
	pb "sale/genproto/sale"
	"sale/storage"
)

type WishlistService struct {
	pb.UnimplementedWishlistServer
	logger *slog.Logger
	repo   storage.Istorage
}

func NewWishlistService(logger *slog.Logger, repo storage.Istorage) *WishlistService {
	return &WishlistService{
		logger: logger,
		repo:   repo,
	}
}

func (s WishlistService) CreateWishlist(ctx context.Context, req *pb.CreateWishlistRequest) (*pb.WishlistResponse, error) {
	s.logger.Info("CreateWishlist rpc method is started")
	res, err := s.repo.Wishlist().CreateWishlist(ctx, req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	s.logger.Info("CreateWishlist finished succesfully")
	return res, nil
}

func (s WishlistService) GetWishlist(ctx context.Context, req *pb.GetWishlistRequest) (*pb.GetWishlistResponse, error) {
	s.logger.Info("GetWishlist rpc method is started")
	res, err := s.repo.Wishlist().GetWishlist(ctx, req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	s.logger.Info("GetWishlist finished succesfully")
	return res, nil
}

func (s WishlistService) GetWishlistById(ctx context.Context, req *pb.GetWishlistByIdRequest) (*pb.GetWishlistByIdResponse, error) {
	s.logger.Info("GetWishlistById rpc method is started")
	res, err := s.repo.Wishlist().GetWishlistById(ctx, req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	s.logger.Info("GetWishlistById finished succesfully")
	return res, nil
}
