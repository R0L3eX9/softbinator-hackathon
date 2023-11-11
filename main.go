package main

import (
	"log"

	"github.com/R0L3eX9/softbinator-hackathon/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func env_init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	env_init()

	router := gin.Default()
    router.GET("/", handlers.Home)
	router.GET("/api/v1/categories", handlers.GetCategories)
	router.POST("/api/v1/create/roadmap", handlers.CreateRoadmap)
    router.POST("/api/v1/ask", handlers.AskGPT)

	err := router.Run()
	if err != nil {
		log.Fatal("ERROR: can not run the server")
		return
	}
}
