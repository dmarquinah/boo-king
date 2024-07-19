package service

import (
	"github.com/dmarquinah/boo-king/internal/app/core/domain/dto"
	"github.com/dmarquinah/boo-king/internal/app/core/domain/entity"
)

type IUserServicePort interface {
	FindUsers() (*[]dto.FindUserDTO, error)
	FindUser(id string) (*dto.FindUserDTO, error)
	CreateUser(user entity.User) (*dto.FindUserDTO, error)
}
