package main

import (
	"log"
	"os"

    "github.com/gin-gonic/gin"
	"github.com/R0L3eX9/softbinator-hackathon/handlers"
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
	API_KEY := os.Getenv("GPT_API_KEY")
	PORT := os.Getenv("PORT")

    router := gin.Default()
    router.GET("/api/v1/categories", handlers.GetCategories)
    // router.POST("/api/create/category")
    // router.POST("/api/create/roadmap")

    err := router.Run("")
    if err != nil {
        log.Fatal("ERROR: can not run the server because {err}")
        return
    }
    log.Println("Server running at: https://localhost:{PORT}")
}
