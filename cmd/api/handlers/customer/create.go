package customer

import (
	"net/http"

	"github.com/dmarquinah/boo-king/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h Handler) CreateCustomer(ctx *gin.Context) {
	// =================> Interpetation
	var customerCreateParams domain.Customer
	if err := ctx.BindJSON(&customerCreateParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// =================> Service

	insertedId, err := h.CustomerService.Create(customerCreateParams)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred..."})
	}

	// =================> Response

	ctx.JSON(http.StatusCreated, gin.H{"player_id": insertedId})
}
