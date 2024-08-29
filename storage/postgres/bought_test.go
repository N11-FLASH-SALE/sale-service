package postgres

import (
	"context"
	pb "sale/genproto/sale"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBought(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := NewBoughtRepository(db)
	req := pb.CreateBoughtRequest{
		UserId:        "246f23a0-81ed-43d2-bf36-931bc38d298c",
		ProductId:     "product456",
		Amount:        2,
		CardNumber:    "1234-5678-9012-3456",
		AmountOfMoney: 10.00,
	}

	ctx := context.Background()
	resp, err := repo.CreateBought(ctx, &req)
	if err != nil {
		t.Error(err)
	}

	// assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Id, "Created Bought ID should not be empty")
}
func TestGetBought(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := NewBoughtRepository(db)
	req := pb.GetBoughtRequest{
		ProductId: "product456",
	}

	ctx := context.Background()
	resp, err := repo.GetBought(ctx, &req)
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, resp)
	assert.Greater(t, len(resp.Boughts), 0, "There should be at least one bought record")
}
func TestGetBoughtOfUser(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := NewBoughtRepository(db)
	req := pb.GetBoughtOfUserRequest{
		UserId: "246f23a0-81ed-43d2-bf36-931bc38d298c",
	}

	ctx := context.Background()
	resp, err := repo.GetBoughtOfUser(ctx, &req)
	if err != nil {
		t.Error(err)
	}

	// assert.NotNil(t, resp)
	assert.Greater(t, len(resp.Boughts), 0, "There should be at least one bought record for the user")
}
