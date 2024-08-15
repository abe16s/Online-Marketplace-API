package models

import (
	"github.com/google/uuid"
)

type Order struct {
    ID        uuid.UUID 		`bson:"_id,omitempty" json:"id,omitempty"`
    BuyerID   uuid.UUID 		`bson:"buyer_id,omitempty" binding:"required" json:"buyer_id,omitempty"`
    ProductID uuid.UUID 		`bson:"product_id,omitempty" binding:"required" json:"product_id,omitempty"`
    Quantity  int      			`bson:"quantity,omitempty" binding:"required" json:"quantity,omitempty"`
    Status    string    		`bson:"status,omitempty" json:"status,omitempty"`
}
