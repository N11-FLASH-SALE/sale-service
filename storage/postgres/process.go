package postgres

import (
	"context"
	"database/sql"
	"time"

	"log/slog"
	pb "sale/genproto/sale"
	logger "sale/logs"
	"sale/storage/repo"
)

type ProcessRepository struct {
	Db *sql.DB
	lg *slog.Logger
}

func NewProcessRepository(db *sql.DB) repo.Processes {
	return &ProcessRepository{Db: db, lg: logger.NewLogger()}
}

func (repo *ProcessRepository) CreateProcess(ctx context.Context, req *pb.CreateProcessRequest) (*pb.ProcessResponse, error) {
	var response pb.ProcessResponse
	query := `INSERT INTO process (user_id, product_id, process_status, amount, created_at)
			  VALUES ($1, $2, $3, $4, $5)
			  RETURNING id;`
	err := repo.Db.QueryRow(query, req.UserId, req.ProductId, req.Status, req.Amount, time.Now().UTC()).
		Scan(&response.Id)

	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (repo *ProcessRepository) GetProcessOfUserByProductId(ctx context.Context, req *pb.GetProcessOfUserByProductIdRequest) (*pb.GetProcessOfUserByProductIdResponse, error) {
	var response pb.GetProcessOfUserByProductIdResponse
	query := `SELECT id, user_id, product_id, process_status, amount
			  FROM process
			  WHERE user_id = $1 AND product_id = $2;`
	rows, err := repo.Db.Query(query, req.UserId, req.ProductId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var process pb.Processes
		err := rows.Scan(&process.Id, &process.UserId, &process.ProductId, &process.Status, &process.Amount)
		if err != nil {
			return nil, err
		}
		response.Processes = append(response.Processes, &process)
	}

	return &response, nil
}

func (repo *ProcessRepository) GetProcessByProductId(ctx context.Context, req *pb.GetProcessByProductIdRequest) (*pb.GetProcessByProductIdResponse, error) {
	var response pb.GetProcessByProductIdResponse
	query := `SELECT id, user_id, product_id, process_status, amount
			  FROM process
			  WHERE product_id = $1;`
	rows, err := repo.Db.Query(query, req.ProductId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var process pb.Processes
		err := rows.Scan(&process.Id, &process.UserId, &process.ProductId, &process.Status, &process.Amount)
		if err != nil {
			return nil, err
		}
		response.Processes = append(response.Processes, &process)
	}

	return &response, nil
}

func (repo *ProcessRepository) UpdateProcess(ctx context.Context, req *pb.UpdateProcessRequest) error {
	query := `UPDATE process
			  SET process_status = $1, updated_at = $2
			  WHERE id = $3;`
	result, err := repo.Db.Exec(query, req.Status, time.Now().UTC(), req.Id)
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (repo *ProcessRepository) CancelProcess(ctx context.Context, req *pb.CancelProcessRequest) error {
	query := `UPDATE process
	SET process_status = 'Cancelled', updated_at = current_timestamp
	WHERE id = $2 and process_status = 'Pending'`
	result, err := repo.Db.Exec(query, req.Id)
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return sql.ErrNoRows
	}
	return nil
}
