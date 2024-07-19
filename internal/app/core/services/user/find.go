package user

import (
	"fmt"
	"log"

	"github.com/dmarquinah/boo-king/internal/app/core/domain/dto"
)

func (s UserService) FindUser(id string) (userDTO *dto.FindUserDTO, err error) {
	result, err := s.UserRepository.FindOne(id)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error obtaining user: %w", err)
	}

	return result, err
}

func (s UserService) FindUsers() (usersDTO *[]dto.FindUserDTO, err error) {
	result, err := s.UserRepository.FindAll()
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error obtaining users: %w", err)
	}

	return result, err
}
