package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sajjad3k/routes"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func main() {
	// Load the env here before calling
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Definir o modo de produção para o Gin Framework
	gin.SetMode(gin.ReleaseMode)

	// Set up the Gin router
	server := os.Getenv("SERVER_ADDRESS")
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Content-type"},
		AllowAllOrigins: true,
	}))

	// Set up the routes
	routes.SetRoutes(router)

	log.Printf("Servidor rodando na porta %s", server)

	// Start the server
	logrus.Error(router.Run(server))
}
