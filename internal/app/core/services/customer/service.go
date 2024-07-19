package customer

import (
	"github.com/dmarquinah/boo-king/internal/app/core/ports/output/repository/customer"
)

type Service struct {
	CustomerRepository customer.ICustomerRepositoryPort
}
