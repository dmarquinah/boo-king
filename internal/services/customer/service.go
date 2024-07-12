package customer

import "github.com/dmarquinah/boo-king/internal/ports"

type Service struct {
	CustomerRepository ports.CustomerRepository
}
