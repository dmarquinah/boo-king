package customer

import (
	"fmt"
	"log"

	"github.com/dmarquinah/boo-king/internal/app/core/domain/entity"
)

func (s Service) Create(customer entity.Customer) (id interface{}, err error) {

	// Set creation time
	//now := time.Now().UTC()
	//customer.CreatedAt = &now

	insertedId, err := s.CustomerRepository.Insert(customer)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error creating customer: %w", err)
	}

	return insertedId, err
}
