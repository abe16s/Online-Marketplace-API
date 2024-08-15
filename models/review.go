package models

import (
	"github.com/google/uuid"
)

type Review struct {
    ID        uuid.UUID 	`bson:"_id,omitempty" binding:"required" json:"id,omitempty"`
    ProductID uuid.UUID 	`bson:"product_id,omitempty" binding:"required" json:"product_id,omitempty"`
    BuyerID   uuid.UUID 	`bson:"buyer_id,omitempty" binding:"required" json:"buyer_id,omitempty"`
    Rating    int       	`bson:"rating,omitempty" binding:"required" json:"rating,omitempty"`
    Comment   string    	`bson:"comment,omitempty" json:"comment,omitempty"`
}
