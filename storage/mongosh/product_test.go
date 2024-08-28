package mongos

import (
	"context"
	pb "sale/genproto/sale"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	db, err := Connect(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	repo := NewProductsRepository(db)
	req := &pb.CreateProductRequest{
		Name:              "Test Product",
		Description:       "This is a test product",
		PriceWithoutStock: 100,
		Stock:             100,
		LimitOfProduct:    1000,
		StartDate:         "2022-01-01",
		EndDate:           "2025-12-31",
		SellerId:          "seller123",
	}
	id, err := repo.CreateProduct(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Product created with ID:", id.Id)
}
