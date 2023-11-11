package handlers

import (
    "log"

	"github.com/R0L3eX9/softbinator-hackathon/mongodb"
	. "github.com/R0L3eX9/softbinator-hackathon/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRoadmap(c *gin.Context) {
    var reqBody []Category
    err := c.Bind(&reqBody)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H {
            "error": "ERROR: couldn't parse request body",
        })
        log.Println(err)
        return
    }
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
