package postgres

import (
	"context"
	"database/sql"
	pb "sale/genproto/sale"
	"sale/storage/repo"
)

type FeedbackRepository struct {
	Db *sql.DB
}

func NewFeedbackRepository(db *sql.DB) repo.Feedback {
	return &FeedbackRepository{Db: db}
}

func (repo *FeedbackRepository) CreateFeedback(ctx context.Context, req *pb.CreateFeedbackRequest) (*pb.FeedbackResponse, error) {
	var res pb.FeedbackResponse
	query := `INSERT INTO feedback ( user_id, product_id, rating, description )
			  VALUES ($1, $2, $3, $4)
			  returning id;`
	err := repo.Db.QueryRow(query, req.UserId, req.ProductId, req.Rating, req.Description).
		Scan(&res.Id)

	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (repo *FeedbackRepository) GetFeedback(ctx context.Context, request *pb.GetFeedbackRequest) (*pb.GetFeedbackResponse, error) {
	var response pb.GetFeedbackResponse
	query := `SELECT user_id, rating, description
			  FROM feedback
			  WHERE product_id = $1;`
	rows, err := repo.Db.Query(query, request.ProductId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var totalRating int64
	var count int64

	for rows.Next() {
		var userId string
		var rating int64
		var description string
		err := rows.Scan(&userId, &rating, &description)
		if err != nil {
			return nil, err
		}
		response.Feedbacks = append(response.Feedbacks, &pb.FeedbackOfProduct{
			UserId:      userId,
			Rating:      rating,
			Description: description,
		})
		totalRating += rating
		count++
	}

	if count > 0 {
		response.AverageRating = totalRating / count
	}

	return &response, nil
}

func (repo *FeedbackRepository) GetFeedbackOfUser(ctx context.Context, request *pb.GetFeedbackOfUserRequest) (*pb.GetFeedbackOfUserResponse, error) {
	var response pb.GetFeedbackOfUserResponse
	query := `SELECT product_id, rating, description
			  FROM feedback
			  WHERE user_id = $1;`
	rows, err := repo.Db.Query(query, request.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product_id string
		var rating int64
		var description string
		err := rows.Scan(&product_id, &rating, &description)
		if err != nil {
			return nil, err
		}
		response.Feedbacks = append(response.Feedbacks, &pb.FeedbackOfUser{
			ProductId:   product_id,
			Rating:      rating,
			Description: description,
		})
	}

	return &response, nil
}
