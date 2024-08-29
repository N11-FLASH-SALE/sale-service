package postgres

import (
	"context"
	pb "sale/genproto/sale"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFeedback(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := NewFeedbackRepository(db)
	req := pb.CreateFeedbackRequest{
		UserId:      "246f23a0-81ed-43d2-bf36-931bc38d298c",
		ProductId:   "product456",
		Rating:      5,
		Description: "Great product!",
	}

	ctx := context.Background()
	resp, err := repo.CreateFeedback(ctx, &req)
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Id, "Created Feedback ID should not be empty")
}
func TestGetFeedback(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := NewFeedbackRepository(db)
	req := pb.GetFeedbackRequest{
		ProductId: "product456",
	}

	ctx := context.Background()
	resp, err := repo.GetFeedback(ctx, &req)
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, resp)
	assert.Greater(t, len(resp.Feedbacks), 0, "There should be at least one feedback record")
	assert.GreaterOrEqual(t, resp.AverageRating, int64(0), "Average rating should be non-negative")
}
func TestGetFeedbackOfUser(t *testing.T) {
	db, err := ConnectionDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := NewFeedbackRepository(db)
	req := pb.GetFeedbackOfUserRequest{
		UserId: "246f23a0-81ed-43d2-bf36-931bc38d298c",
	}

	ctx := context.Background()
	resp, err := repo.GetFeedbackOfUser(ctx, &req)
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, resp)
	assert.Greater(t, len(resp.Feedbacks), 0, "There should be at least one feedback record for the user")
}
