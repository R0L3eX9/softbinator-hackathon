package models

import "fmt"

type Roadmap struct {
    ID string `json:"id"`
    Title string `json:"title"`
    Difficulties []Difficulty `json:"difficulties"`
}

type Difficulty struct {
    Level string `json:"level"`
    Skills []Skill `json:"skills"`
}

type Skill struct {
    Title string `json:"title"`
    ToLearn []string `json:"tolearn"`
    Resources []Resource `json:"resources"`
}

type Resource struct {
    Link string `json:"link"`
    Description string `json:"description"`
}

func Test() {
    ex1 := Roadmap {
        ID: "fdafk",
        Title: "Python",
        Difficulties: []Difficulty{
            Difficulty {
                Level: "Easy",
                Skills: []Skill {
                    Skill {
                        Title: "Intro",
                        ToLearn: []string {
                            "variables",
                            "operators",
                        },
                        Resources: []Resource {
                            Resource {
                                Link: "https://test.com",
                                Description: "test",
                            },
                            Resource {
                                Link: "https://test2.com",
                                Description: "test2",
                            },
                        },
                    },
                    Skill {
                        Title: "Control flow",
                        ToLearn: []string {
                            "variables",
                            "operators",
                        },
                        Resources: []Resource {
                            Resource {
                                Link: "https://test.com",
                                Description: "test",
                            },
                            Resource {
                                Link: "https://test2.com",
                                Description: "test2",
                            },
                        },
                    },
                },
            },
        },
    }
    fmt.Println(ex1)
}
