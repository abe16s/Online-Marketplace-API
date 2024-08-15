package models

import (
	"github.com/google/uuid"
)

type User struct {
    ID       uuid.UUID 		`bson:"_id" json:"id,omitempty"`
    Name     string     	`bson:"name,omitempty" json:"name,omitempty"`
    Email    string         `bson:"email" binding:"required" json:"email" validate:"required,email"`
    Password string     	`bson:"password" binding:"required" json:"password" validate:"required,min=8,password"`
    IsSeller bool     	    `bson:"is_seller,omitempty" json:"is_seller,omitempty"`
}


