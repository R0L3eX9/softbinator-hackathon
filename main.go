package main

import (
    "log"
    "os"
    "fmt"

    "github.com/joho/godotenv"
    "github.com/R0L3eX9/softbinator-hackathon/models"
//    "github.com/R0L3eX9/softbinator-hackathon/handlers"
    // "net/http"
    // "github.com/gin-gonic/gin"
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
    fmt.Println(API_KEY)
    models.Test()
    // handlers.Test()
}
