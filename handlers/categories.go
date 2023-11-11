package handlers

import (
    "log"
    "os"
	"net/http"
    "bytes"
    "encoding/json"

	"github.com/R0L3eX9/softbinator-hackathon/mongodb"
	. "github.com/R0L3eX9/softbinator-hackathon/models"
	"github.com/gin-gonic/gin"
)

const openAIURL = "https://api.openai.com/v1/engines/davinci-codex/completions"

func AskGPT(c *gin.Context) {
    var req UserPrompt
    if err := c.BindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    gptMsg := GPTMessage {
        Role: "system",
        Content: GPT_PROMPT + req.Keywords,
    }

    gptPrompt := GPTPrompt {
        Model: "gpt-3.5-turbo",
        Message: gptMsg,
        MaxTokens: 1000,
    }

    body, err := json.Marshal(gptPrompt)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    //
    resp, err := http.Post(openAIURL, "application/json", bytes.NewBuffer(body))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer resp.Body.Close()
    //
    API_KEY := os.Getenv("GPT_API_KEY")
    c.Request.Header.Set("Authorization", "Bearer " + API_KEY)
    var result map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, result)
}

func CreateRoadmap(c *gin.Context) {
    var reqBody Roadmap
    err := c.Bind(&reqBody)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H {
            "error": "ERROR: couldn't parse request body",
        })
        log.Println(err)
        return
    }
    mongodb.AddUserRoadmap(reqBody)
    c.JSON(http.StatusOK, reqBody)
}

func GetCategories(c *gin.Context) {
    data, err := mongodb.DBRead()
    if err != nil {
        log.Println(err)
        return
    }
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H {
            "error": "ERROR: couldn't parse request body",
        })
        log.Println(err)
        return
    }
    c.JSON(http.StatusOK, data)
}

func Home(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H {
        "message": "Hello world",
    })
}
