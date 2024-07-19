package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dmarquinah/boo-king/cmd/api/core"

	userRouter "github.com/dmarquinah/boo-king/internal/app/adapters/input/http/user"
	"github.com/dmarquinah/boo-king/internal/app/adapters/output/repository/mongo"
	userRepository "github.com/dmarquinah/boo-king/internal/app/adapters/output/repository/mongo/user"
	userService "github.com/dmarquinah/boo-king/internal/app/core/services/user"
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

	database := client.Database(os.Getenv("DB_NAME"))

	userRepo := userRepository.Repository{
		Database: database,
	}

	userSrv := userService.UserService{
		UserRepository: userRepo,
	}

	userRouter := userRouter.UserRouter{
		UserService: userSrv,
	}

	server.GinEngine.GET("/user/:id", userRouter.GetUser)
	server.GinEngine.POST("/user", userRouter.CreateUser)

	log.Fatalln(server.GinEngine.Run(":8001"))
}
