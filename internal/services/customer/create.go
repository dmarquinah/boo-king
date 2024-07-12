package customer

import (
	"fmt"
	"log"
	"time"

	"github.com/dmarquinah/boo-king/internal/domain"
)

func (s Service) Create(customer domain.Customer) (id interface{}, err error) {

	// Set creation time
	customer.CreationTime = time.Now().UTC()

	insertedId, err := s.CustomerRepository.Insert(customer)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error creating customer: %w", err)
	}

	return insertedId, err
}
