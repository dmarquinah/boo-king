package user

import (
	"github.com/dmarquinah/boo-king/internal/app/core/domain/dto"
	"github.com/dmarquinah/boo-king/internal/app/core/domain/entity"
)

type IUserRepositoryPort interface {
	FindOne(id string) (userDTO *dto.FindUserDTO, err error)
	FindAll() (usersDTO *[]dto.FindUserDTO, err error)
	Insert(user entity.User) (userDTO *dto.FindUserDTO, err error)
}
