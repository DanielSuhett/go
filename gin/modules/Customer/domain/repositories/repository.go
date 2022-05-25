package customer

import (
	"github.com/DanielSuhett/go/gin/modules/Customer/domain/entities"
)

type CustomerRepository interface {
	Get(id string) (*entities.Customer, error)
	Add(entities.Customer) error
	Update(id string, customer entities.Customer) error
}
