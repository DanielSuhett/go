package services

import (
	"context"
	"log"

	"github.com/DanielSuhett/go/gin/modules/Customer/domain/entities"
	"github.com/DanielSuhett/go/gin/modules/Customer/domain/repositories"
	impl "github.com/DanielSuhett/go/gin/modules/Customer/infra"
	"github.com/google/uuid"
)

type CustomerService struct {
	customers customer.CustomerRepository
}

type Service func(cs *CustomerService) error

func Run() (*CustomerService, error) {
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

	customer := entities.Customer{ID: uuid.NewString(), Name: name, Email: email, Addresses: addresses}

	err := c.customers.Add(customer)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (c *CustomerService) GetCustomer(id string) (*entities.Customer, error) {
	customer, err := c.customers.Get(id)

	if err != nil {
		return customer, err
	}

	return customer, nil
}
