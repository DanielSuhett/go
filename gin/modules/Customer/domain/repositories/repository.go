package customer

import (
	"github.com/DanielSuhett/go/gin/modules/Customer/domain/entities"

	"github.com/google/uuid"
)

type CustomerRepository interface {
	Get(uuid.UUID) (*entities.Customer, error)
	Add(entities.Customer) error
	Update(uuid.UUID, entities.Customer) error
}