package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	. "github.com/R0L3eX9/softbinator-hackathon/models"
	"github.com/R0L3eX9/softbinator-hackathon/mongodb"
	"github.com/gin-gonic/gin"
)

const openAIURL = "https://api.openai.com/v1/chat/completions"

func AskGPT(c *gin.Context) {
	var req UserPrompt
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gptMsg := []GPTMessage{
		{
			Role:    "system",
			Content: "You are an assistant!",
		},
		{
			Role:    "user",
			Content: GPT_PROMPT + req.Keywords,
		},
	}

	gptPrompt := GPTPrompt{
		Model:     "gpt-3.5-turbo",
		Messages:  gptMsg,
		MaxTokens: 1000,
	}

	body, err := json.Marshal(gptPrompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	openaiReq, err := http.NewRequest("POST", openAIURL, bytes.NewBuffer(body))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	API_KEY := os.Getenv("GPT_API_KEY")
	openaiReq.Header.Set("Content-Type", "application/json")
	openaiReq.Header.Set("Authorization", "Bearer "+API_KEY)

	client := &http.Client{}
	resp, err := client.Do(openaiReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

    var result map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, result)

    apiChoices, ok := result["choices"].([]interface{});
    if !ok {
        log.Println("Choices failed not found")
        return
    }

    apiChoice, ok := apiChoices[0].(map[string]interface{})
    if !ok {
        log.Println("Choice failed not found")
        return
    }

    apiMessage, ok := apiChoice["message"].(map[string]interface{})
    if !ok {
        log.Println("Message failed not found")
        return
    }
    apiResult := apiMessage["content"]

    customRoadmap, ok := apiResult.(Roadmap)
    if !ok {
        log.Println("Couldn't parse into custom struct")
        return
    }
    log.Println(customRoadmap.Name)

    err = mongodb.AddUserRoadmap(customRoadmap)
    if err != nil {
        log.Println(err)
        return
    }
    c.JSON(http.StatusOK, customRoadmap)
}

func CreateRoadmap(c *gin.Context) {
	var reqBody Roadmap
	err := c.Bind(&reqBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "ERROR: couldn't parse request body",
		})
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, data)
}

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello world",
	})
}
