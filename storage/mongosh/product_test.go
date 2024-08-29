package mongosh

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

func TestGetProduct(t *testing.T) {
	db, err := Connect(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	repo := NewProductsRepository(db)
	req := &pb.GetProductRequest{
		MaxPrice: 1000,
	}
	resp, err := repo.GetProduct(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Product retrieved:", resp)
}

func TestGetProductById(t *testing.T) {
	db, err := Connect(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	repo := NewProductsRepository(db)
	req := &pb.ProductId{
		Id: "66cf7076533ee98a300b9020",
	}
	resp, err := repo.GetProductById(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Product retrieved by ID:", resp)
}

func TestUpdateProduct(t *testing.T) {
	db, err := Connect(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	repo := NewProductsRepository(db)
	req := &pb.UpdateProductRequest{
		Id:      "66cf7076533ee98a300b9020",
		Color:   []string{"red", "yellow"},
		EndDate: "2024-08-29",
	}
	err = repo.UpdateProduct(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Product updated successfully")
}

func TestDeleteProduct(t *testing.T) {
	db, err := Connect(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	repo := NewProductsRepository(db)
	req := &pb.ProductId{
		Id: "66cf7055930e6a2ff0f197a7",
	}
	err = repo.DeleteProduct(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Product deleted successfully")
}

func TestIsProductOk(t *testing.T) {
	db, err := Connect(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	repo := NewProductsRepository(db)
	req := &pb.ProductId{
		Id: "66cf7076533ee98a300b9020",
	}
	err = repo.IsProductOk(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Product is OK")
}

func TestAddPhotosToProduct(t *testing.T) {
	db, err := Connect(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	repo := NewProductsRepository(db)
	req := &pb.AddPhotosRequest{
		ProductId: "66cf7076533ee98a300b9020",
		PhotoUrl:  "photo1.jpg",
	}
	err = repo.AddPhotosToProduct(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Photos added to product successfully")
}

func TestDeletePhotosFromProduct(t *testing.T) {
	db, err := Connect(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	repo := NewProductsRepository(db)
	req := &pb.DeletePhotosRequest{
		ProductId: "66cf7076533ee98a300b9020",
		PhotoUrl:  "photo2.jpg",
	}
	err = repo.DeletePhotosFromProduct(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Photos deleted from product successfully")
}
