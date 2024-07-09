package main

import (
	"log"
	"net/http"

	"github.com/dmarquinah/boo-king/cmd/api/core"
	"github.com/gin-gonic/gin"
)

func main() {
	server := core.Server{GinEngine: gin.Default()}

	server.GinEngine.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong!"})
	})

	log.Fatalln(server.GinEngine.Run(":8001"))
}
