package models

import (
	"github.com/google/uuid"
)

type Product struct {
    ID          uuid.UUID 		`bson:"_id,omitempty" binding:"required" json:"id,omitempty"`
    Name        string		    `bson:"name,omitempty" json:"name,omitempty"`
    Description string      	`bson:"description,omitempty" json:"description,omitempty"`
    Price       float64     	`bson:"price,omitempty" binding:"required" json:"price,omitempty"`
    SellerID    uuid.UUID 		`bson:"seller_id,omitempty" binding:"required" json:"seller_id,omitempty"`
    ImageURL    string      	`bson:"image_url,omitempty" json:"image_url,omitempty"`
}
