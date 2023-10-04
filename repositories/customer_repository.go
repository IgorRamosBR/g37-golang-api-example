package repositories

import (
	"errors"
	"fmt"
	"postech/models"
)

type CustomerRepository interface {
	GetAllCustomers() []string
	SaveCustomer(customer models.Customer) error
}

type customerRepository struct {
}

func NewCustomerRepository() CustomerRepository {
	return customerRepository{}
}

func (r customerRepository) GetAllCustomers() []string {
	return []string{
		"Igor",
		"Alvaro",
		"Bruno",
		"Gildo",
	}
}

func (r customerRepository) SaveCustomer(customer models.Customer) error {
	fmt.Print(customer.Name)
	if customer.Name == "" {
		return errors.New("name is empty")
	}

	return nil
}
