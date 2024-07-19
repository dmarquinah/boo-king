package user

import (
	"net/http"

	"github.com/dmarquinah/boo-king/internal/app/core/domain/entity"
	"github.com/gin-gonic/gin"
)

func (r UserRouter) CreateUser(ctx *gin.Context) {
	// =================> Interpetation
	var userCreateParams entity.User
	if err := ctx.BindJSON(&userCreateParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// =================> Service

	user, err := r.UserService.CreateUser(userCreateParams)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred..."})
	}

	// =================> Response

	ctx.JSON(http.StatusCreated, user)
}
