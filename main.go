package main

import (
	"log"
	"holamundo/core"
	"holamundo/launch"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	core.InitializeApp()

	
	app := gin.Default()
	app.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:8080", "http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	launch.RegisterRoutes(app)

	log.Println("API corriendo en http://localhost:8080")
	if err := app.Run(":8080"); err != nil {
		log.Fatalf("Error al correr el servidor: %v", err)
	}
}
