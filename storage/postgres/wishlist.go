package postgres

import (
	"context"
	"database/sql"
	"log/slog"
	pb "sale/genproto/sale"
	logger "sale/logs"
	"sale/storage/repo"
	"time"
)

type WishlistRepository struct {
	Db *sql.DB
	lg *slog.Logger
}

func NewWishlistRepository(db *sql.DB) repo.Wishlist {
	return &WishlistRepository{Db: db, lg: logger.NewLogger()}
}

func (repo *WishlistRepository) CreateWishlist(ctx context.Context, request *pb.CreateWishlistRequest) (*pb.WishlistResponse, error) {
	response := pb.WishlistResponse{ProductId: request.ProductId}
	query := `INSERT INTO wishlist (user_id, product_id, created_at)
			  VALUES ($1, $2, $3)
			  RETURNING id;`
	err := repo.Db.QueryRowContext(ctx, query, request.UserId, request.ProductId, time.Now().UTC()).
		Scan(&response.Id)

	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (repo *WishlistRepository) GetWishlist(ctx context.Context, request *pb.GetWishlistRequest) (*pb.GetWishlistResponse, error) {
	var response pb.GetWishlistResponse
	query := `SELECT id, product_id
			  FROM wishlist
			  WHERE user_id = $1;`
	rows, err := repo.Db.QueryContext(ctx, query, request.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		wishes := pb.WishlistResponse{}
		err := rows.Scan(&wishes.Id, &wishes.ProductId)
		if err != nil {
			return nil, err
		}
		response.Wishes = append(response.Wishes, &wishes)
	}

	return &response, nil
}

func (repo *WishlistRepository) GetWishlistById(ctx context.Context, request *pb.GetWishlistByIdRequest) (*pb.GetWishlistByIdResponse, error) {
	var response pb.GetWishlistByIdResponse
	query := `SELECT product_id
			  FROM wishlist
			  WHERE id = $1
			  LIMIT 1;`
	err := repo.Db.QueryRowContext(ctx, query, request.Id).Scan(&response.ProductId)

	if err != nil {
		return nil, err
	}
	return &response, nil
}
