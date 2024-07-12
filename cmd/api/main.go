package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dmarquinah/boo-king/cmd/api/core"
	"github.com/dmarquinah/boo-king/cmd/api/handlers/customer"

	"github.com/dmarquinah/boo-king/internal/repositories/mongo"
	customerRepository "github.com/dmarquinah/boo-king/internal/repositories/mongo/customer"
	customerService "github.com/dmarquinah/boo-king/internal/services/customer"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := core.Server{GinEngine: gin.Default()}

	server.GinEngine.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong!"})
	})

	client, err := mongo.ConnectClient(os.Getenv("MONGO_URI"))
	if err != nil {
		log.Fatal(err.Error())
	}

	customerRepo := customerRepository.Repository{
		Client: client,
	}

	customerSrv := customerService.Service{
		CustomerRepository: customerRepo,
	}

	customerHandler := customer.Handler{
		CustomerService: customerSrv,
	}

	server.GinEngine.POST("/customers", customerHandler.CreateCustomer)

	log.Fatalln(server.GinEngine.Run(":8001"))
}
