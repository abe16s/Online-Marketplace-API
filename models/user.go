package models

import (
	"github.com/google/uuid"
)

type User struct {
    ID       uuid.UUID 		`bson:"_id" binding:"required" json:"id,omitempty"`
    Name     string     	`bson:"name,omitempty" json:"name,omitempty"`
    Email    string     	`bson:"email" binding:"required" json:"email"`
    Password string     	`bson:"password" binding:"required" json:"password"`
    IsSeller string     	`bson:"is_seller,omitempty" json:"is_seller,omitempty"`
}


