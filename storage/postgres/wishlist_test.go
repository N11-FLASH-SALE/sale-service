package postgres

import (
	"fmt"
	"testing"

	pb "sale/genproto/sale"

	"github.com/stretchr/testify/assert"
)

func TestCreateWishlist(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	repo := NewWishlistRepository(db)

	req := &pb.CreateWishlistRequest{
		UserId:    "246f23a0-81ed-43d2-bf36-931bc38d298c",
		ProductId: "product456",
	}

	resp, err := repo.CreateWishlist(req)
	if err != nil {
		t.Fatalf("Failed to create wishlist: %v", err)
	}

	assert.NotEmpty(t, resp.Id)
	fmt.Printf("Created Wishlist ID: %s\n", resp.Id)
}

func TestGetWishlist(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	repo := NewWishlistRepository(db)

	req := &pb.GetWishlistRequest{
		UserId: "246f23a0-81ed-43d2-bf36-931bc38d298c",
	}

	resp, err := repo.GetWishlist(req)
	if err != nil {
		t.Fatalf("Failed to get wishlist: %v", err)
	}

	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.ProductId)
	fmt.Printf("Found %d products in wishlist for user %s\n", len(resp.ProductId), req.UserId)
}

func TestGetWishlistById(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	repo := NewWishlistRepository(db)

	req := &pb.GetWishlistRequest{
		UserId: "246f23a0-81ed-43d2-bf36-931bc38d298c",
	}

	resp, err := repo.GetWishlistById(req)
	if err != nil {
		t.Fatalf("Failed to get wishlist by ID: %v", err)
	}

	assert.NotEmpty(t, resp.Id)
	fmt.Printf("Found wishlist with ID %s for user %s\n", resp.Id, req.UserId)
}
