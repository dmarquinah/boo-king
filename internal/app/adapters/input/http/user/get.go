package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r UserRouter) GetUser(ctx *gin.Context) {
	// =================> Interpetation
	userIdParam := ctx.Param("id")

	// =================> Service

	user, err := r.UserService.FindUser(userIdParam)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred..."})
		return
	}

	// =================> Response

	ctx.JSON(http.StatusOK, user)
}
