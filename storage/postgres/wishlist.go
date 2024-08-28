package postgres

import (
	"database/sql"
	"log/slog"
	logger "sale/logs"
	pb "sale/genproto/sale"
	"time"
)

type WishlistRepository struct {
	Db *sql.DB
	lg *slog.Logger
}

func NewWishlistRepository(db *sql.DB) *WishlistRepository {
	return &WishlistRepository{Db: db, lg: logger.NewLogger()}
}

func (repo *WishlistRepository) CreateWishlist(request *pb.CreateWishlistRequest) (*pb.WishlistResponse, error) {
	var response pb.WishlistResponse
	query := `INSERT INTO wishlist (user_id, product_id, created_at)
			  VALUES ($1, $2, $3)
			  RETURNING id;`
	err := repo.Db.QueryRow(query, request.UserId, request.ProductId, time.Now().UTC()).
		Scan(&response.Id)

	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (repo *WishlistRepository) GetWishlist(request *pb.GetWishlistRequest) (*pb.GetWishlistResponse, error) {
	var response pb.GetWishlistResponse
	query := `SELECT product_id
			  FROM wishlist
			  WHERE user_id = $1;`
	rows, err := repo.Db.Query(query, request.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var productId string
		err := rows.Scan(&productId)
		if err != nil {
			return nil, err
		}
		response.ProductId = append(response.ProductId, productId)
	}

	return &response, nil
}

func (repo *WishlistRepository) GetWishlistById(request *pb.GetWishlistRequest) (*pb.WishlistResponse, error) {
	var response pb.WishlistResponse
	query := `SELECT id
			  FROM wishlist
			  WHERE user_id = $1
			  LIMIT 1;`
	err := repo.Db.QueryRow(query, request.UserId).
		Scan(&response.Id)

	if err != nil {
		return nil, err
	}
	return &response, nil
}
