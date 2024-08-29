package postgres

import (
	"context"
	"database/sql"

	pb "sale/genproto/sale"
	"sale/storage/repo"
)

type BoughtRepository struct {
	db *sql.DB
}

func NewBoughtRepository(db *sql.DB) repo.Bought {
	return &BoughtRepository{db: db}
}

func (r *BoughtRepository) CreateBought(ctx context.Context, req *pb.CreateBoughtRequest) (*pb.BoughtResponse, error) {
	var res pb.BoughtResponse
	query := `
		INSERT INTO bought (user_id, product_id, amount, card_number, amount_of_money)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`
	err := r.db.QueryRowContext(ctx, query, req.UserId, req.ProductId, req.Amount, req.CardNumber, req.AmountOfMoney).
		Scan(&res.Id)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *BoughtRepository) GetBought(ctx context.Context, req *pb.GetBoughtRequest) (*pb.GetBoughtResponse, error) {
	var res pb.GetBoughtResponse
	query := `
		SELECT user_id, amount, card_number, amount_of_money
		FROM bought
		WHERE product_id = $1;
	`
	rows, err := r.db.QueryContext(ctx, query, req.ProductId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var bought pb.BoughtOfProduct
		err := rows.Scan(&bought.UserId, &bought.Amount, &bought.CardNumber, &bought.AmountOfMoney)
		if err != nil {
			return nil, err
		}
		res.Boughts = append(res.Boughts, &bought)
	}

	return &res, nil
}

func (r *BoughtRepository) GetBoughtOfUser(ctx context.Context, req *pb.GetBoughtOfUserRequest) (*pb.GetBoughtOfUserResponse, error) {
	var res pb.GetBoughtOfUserResponse
	query := `
		SELECT product_id, amount, card_number, amount_of_money
		FROM bought
		WHERE user_id = $1;
	`
	rows, err := r.db.QueryContext(ctx, query, req.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var bought pb.BoughtOfProduct
		err := rows.Scan(&bought.UserId, &bought.Amount, &bought.CardNumber, &bought.AmountOfMoney)
		if err != nil {
			return nil, err
		}
		res.Boughts = append(res.Boughts, &bought)
	}

	return &res, nil
}