package ports

import "github.com/dmarquinah/boo-king/internal/domain"

type CustomerService interface {
	Create(customer domain.Customer) (id interface{}, err error)
}

type CustomerRepository interface {
	Insert(customer domain.Customer) (id interface{}, err error)
}
