package entities

import (
	"github.com/google/uuid"
)

type Customer struct {
	ID uuid.UUID
	Name string
	Email string
	Addresses []Address
}