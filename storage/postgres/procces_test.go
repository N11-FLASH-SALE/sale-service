package postgres

import (
	"context"
	"fmt"
	"testing"

	pb "sale/genproto/sale"

	"github.com/stretchr/testify/assert"
)

func TestCreateProcess(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := NewProcessRepository(db)
	req := pb.CreateProcessRequest{
		UserId:    "user123",
		ProductId: "product456",
		Status:    "pending",
		Amount:    1000,
	}

	ctx := context.Background()
	resp, err := repo.CreateProcess(ctx, &req)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Created Process ID:", resp)
}

func TestGetProcessOfUserByProductId(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := NewProcessRepository(db)
	req := pb.GetProcessOfUserByProductIdRequest{
		UserId:    "user123",
		ProductId: "product456",
	}

	ctx := context.Background()
	resp, err := repo.GetProcessOfUserByProductId(ctx,&req)
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Processes)
	for _, process := range resp.Processes {
		assert.Equal(t, req.UserId, process.UserId)
		assert.Equal(t, req.ProductId, process.ProductId)
	}
	fmt.Printf("Found %d processes for user %s and product %s\n", len(resp.Processes), req.UserId, req.ProductId)
}

func TestGetProcessByProductId(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := NewProcessRepository(db)
	req := pb.GetProcessByProductIdRequest{
		ProductId: "product456",
	}

	ctx := context.Background()
	resp, err := repo.GetProcessByProductId(ctx, &req)
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Processes)
	for _, process := range resp.Processes {
		assert.Equal(t, req.ProductId, process.ProductId)
	}
	fmt.Printf("Found %d processes for product %s\n", len(resp.Processes), req.ProductId)
}

func TestUpdateProcess(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := NewProcessRepository(db)
	req := pb.UpdateProcessRequest{
		Id:     "process789",
		Status: "completed",
	}

	ctx := context.Background()
	err = repo.UpdateProcess(ctx, &req)
	assert.NoError(t, err)
	fmt.Printf("Updated process %s status to %s\n", req.Id, req.Status)
}

func TestCancelProcess(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := NewProcessRepository(db)
	req := pb.CancelProcessRequest{
		Id: "process789", // Ensure this ID exists in your test database
	}

	ctx := context.Background()
	err = repo.CancelProcess(ctx, &req)
	assert.NoError(t, err)
	fmt.Printf("Cancelled process %s\n", req.Id)
}
