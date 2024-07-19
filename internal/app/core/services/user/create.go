package user

import (
	"fmt"
	"log"

	"github.com/dmarquinah/boo-king/internal/app/core/domain/dto"
	"github.com/dmarquinah/boo-king/internal/app/core/domain/entity"
)

func (s UserService) CreateUser(user entity.User) (*dto.FindUserDTO, error) {
	result, err := s.UserRepository.Insert(user)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error obtaining user: %w", err)
	}

	return result, err
}
