package user

import "github.com/dmarquinah/boo-king/internal/app/core/ports/output/repository/user"

type UserService struct {
	UserRepository user.IUserRepositoryPort
}
