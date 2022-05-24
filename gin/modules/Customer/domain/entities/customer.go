package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string
	Email     string
	Addresses []Address
}

type CreateCustomer struct {
	Name    string  `form:"name" binding:"required"`
	Email   string  `form:"email" binding:"required"`
	Address Address `form:"address" binding:"omitempty"`
}
