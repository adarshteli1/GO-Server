package main

import (
	"Go-Server/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf("Error whiule loading the env %v\n", err)
		return
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.Staticroutes(router)

	fmt.Println("Server is Starting")

	router.Run(":" + port)

}
