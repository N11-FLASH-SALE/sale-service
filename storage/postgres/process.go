package postgres

import (
	"context"
	"database/sql"

	pb "sale/genproto/sale"
	"sale/storage/repo"
)

type ProcessRepository struct {
	Db *sql.DB
}

func NewProcessRepository(db *sql.DB) repo.Processes {
	return &ProcessRepository{Db: db}
}

func (repo *ProcessRepository) CreateProcess(ctx context.Context, req *pb.CreateProcessRequest) (*pb.ProcessResponse, error) {
	var response pb.ProcessResponse
	query := `INSERT INTO process (user_id, product_id, amount)
			  VALUES ($1, $2, $3)
			  RETURNING id;`
	err := repo.Db.QueryRow(query, req.UserId, req.ProductId, req.Amount).
		Scan(&response.Id)

	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (repo *ProcessRepository) GetProcessOfUserByProductId(ctx context.Context, req *pb.GetProcessOfUserByProductIdRequest) (*pb.GetProcessOfUserByProductIdResponse, error) {
	var response pb.GetProcessOfUserByProductIdResponse
	query := `SELECT id, user_id, product_id, status, amount
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

func (repo *ProcessRepository) GetProcessByUserId(ctx context.Context, req *pb.GetProcessByUserIdRequest) (*pb.GetProcessByUserIdResponse, error) {
	var response pb.GetProcessByUserIdResponse
	query := `SELECT id, user_id, product_id, status, amount
			  FROM process
			  WHERE user_id = $1 AND user_id = $2;`
	rows, err := repo.Db.Query(query, req.UserId, req.UserId)
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
	query := `SELECT id, user_id, product_id, status, amount
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
			  SET status = $1
			  WHERE id = $2;`
	result, err := repo.Db.Exec(query, req.Status, req.Id)
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

func (repo *ProcessRepository) CancelProcess(ctx context.Context, req *pb.CancelProcessRequest) (*pb.CancelProcessResponse, error) {
	var response pb.CancelProcessResponse
	query := `UPDATE process
	SET status = 'Cancelled'
	WHERE id = $2 and status = 'Pending' RETURNING amount`
	err := repo.Db.QueryRowContext(ctx, query, req.Id).Scan(&response.Amount)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (repo *ProcessRepository) GetProcessById(ctx context.Context, req *pb.GetProcessByIdRequest) (*pb.GetProcessByIdResponse, error) {
	var response pb.GetProcessByIdResponse
	query := `SELECT id, user_id, product_id, status, amount
			  FROM process
			  WHERE id = $1;`
	err := repo.Db.QueryRowContext(ctx, query, req.Id).Scan(&response.Id, &response.UserId, &response.ProductId, &response.Status, &response.Amount)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
