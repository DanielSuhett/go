package services

import (
	"context"
	"log"

	"github.com/DanielSuhett/go/gin/modules/Customer/domain/entities"
	"github.com/DanielSuhett/go/gin/modules/Customer/domain/repositories"
	impl "github.com/DanielSuhett/go/gin/modules/Customer/infra"
)

type CustomerService struct {
	customers customer.CustomerRepository
}

type Service func(cs *CustomerService) error

func NewOrderService() (*CustomerService, error) {
	s := &CustomerService{}

	cs, err := impl.Connect(context.Background())
	if err != nil {
		return &CustomerService{}, err
	}

	s.customers = cs

	return s, nil
}



func (c *CustomerService) CreateCustomer(name string, email string, address interface{}) error {
	addr := address.(entities.Address)
	addresses := []entities.Address{addr}

	customer := entities.Customer{Name: name, Email: email, Addresses: addresses}

	err := c.customers.Add(customer)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
