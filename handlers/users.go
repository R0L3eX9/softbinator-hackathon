package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    . "github.com/R0L3eX9/softbinator-hackathon/models"
)

func Test() {
    return
}
func GetCategories(c *gin.Context) {
    sample := Category {
        Name: "Coding",
        Roadmap: Roadmap {
            Title: "Python",
            Difficulties: []Difficulty {
                Difficulty {
                    Level: "EASY",
                    Skills: []Skill {
                        Skill {
                            Title: "Syntax",
                            Description: "Basic python syntax",
                            State: false,
                            Resources: []Resource {
                                Resource {
                                    Name: "Python for noobs",
                                    Description: "Python book for beginners",
                                },
                            },
                        },
                    },
                },
            },
        },
    }
    c.IndentedJSON(http.StatusOK, sample)

}
