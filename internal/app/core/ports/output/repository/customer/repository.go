package customer

import "github.com/dmarquinah/boo-king/internal/app/core/domain/entity"

type ICustomerRepositoryPort interface {
	Insert(customer entity.Customer) (id interface{}, err error)
}
