package ports

import (
	"github.com/dmarquinah/boo-king/internal/app/core/domain/entity"
)

type CustomerService interface {
	Create(customer entity.Customer) (id interface{}, err error)
}

type CustomerRepository interface {
	Insert(customer entity.Customer) (id interface{}, err error)
}
