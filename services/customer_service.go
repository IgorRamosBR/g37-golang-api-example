package services

import (
	"postech/models"
	"postech/repositories"
)

type CustomerService interface {
	GetAllCustomers() []string
	SaveCustomer(customer models.Customer) error
}

type customerService struct {
	repository repositories.CustomerRepository
}

func NewCustomerService(repository repositories.CustomerRepository) CustomerService {
	return customerService{
		repository: repository,
	}
}

func (s customerService) GetAllCustomers() []string {
	return s.repository.GetAllCustomers()
}

func (s customerService) SaveCustomer(customer models.Customer) error {
	return s.repository.SaveCustomer(customer)
}
