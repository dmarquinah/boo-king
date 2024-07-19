package user

import "github.com/dmarquinah/boo-king/internal/app/core/ports/input/service"

type UserRouter struct {
	UserService service.IUserServicePort
}
