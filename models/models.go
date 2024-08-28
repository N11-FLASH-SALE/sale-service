package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Name           string             `bson:"name"`
	Description    string             `bson:"description"`
	Price          float64            `bson:"price"`
	Stock          int64              `bson:"stock"`
	PriceWithStock float64            `bson:"price_with_stock"`
	LimitOfProduct int64              `bson:"limit_of_product"`
	Size           []string           `bson:"size"`
	Color          []string           `bson:"color"`
	StartDate      time.Time          `bson:"start_date"`
	EndDate        time.Time          `bson:"end_date"`
	SellerID       string             `bson:"seller_id"`
}
