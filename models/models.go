package models

type Category struct {
    Name string `json:"categoryName"`
    Roadmap Roadmap `json:"roadmapName"`
}
type Roadmap struct {
    Title string `json:"title"`
    Difficulties []Difficulty `json:"difficulties"` }

type Difficulty struct {
    Level string `json:"difficultyLevel"`
    Skills []Skill `json:"skills"`
}

type Skill struct {
    Title string `json:"title"`
    Description string `json:"description"`
    State bool `json:"state"`
    Resources []Resource `json:"resources"`
}

type Resource struct {
    Name string `json:"resourceName"`
    Description string `json:"description"`
}
