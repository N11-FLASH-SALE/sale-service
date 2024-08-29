package postgres

import (
	"context"
	pb "sale/genproto/sale"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateWishlist(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := NewWishlistRepository(db)
	req := pb.CreateWishlistRequest{
		UserId:    "246f23a0-81ed-43d2-bf36-931bc38d298c",
		ProductId: "product456",
	}

	ctx := context.Background()
	resp, err := repo.CreateWishlist(ctx, &req)
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Id, "Created Wishlist ID should not be empty")
}
func TestGetWishlist(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := NewWishlistRepository(db)
	req := pb.GetWishlistRequest{
		UserId: "246f23a0-81ed-43d2-bf36-931bc38d298c",
	}

	ctx := context.Background()
	resp, err := repo.GetWishlist(ctx, &req)
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, resp)
	assert.Greater(t, len(resp.Wishes), 0, "There should be at least one wishlist record")
}
func TestGetWishlistById(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := NewWishlistRepository(db)
	req := pb.GetWishlistByIdRequest{
		Id: "b062c900-3e59-43ae-8a07-b3e12a3db4fe",
	}

	ctx := context.Background()
	resp, err := repo.GetWishlistById(ctx, &req)
	if err != nil {
		t.Error(err)
	}

	// assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.ProductId, "Wishlist Product ID should not be empty")
}
