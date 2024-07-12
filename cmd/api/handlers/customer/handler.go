package customer

import "github.com/dmarquinah/boo-king/internal/ports"

type Handler struct {
	CustomerService ports.CustomerService
}
