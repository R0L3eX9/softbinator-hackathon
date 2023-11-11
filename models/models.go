package models

type Category struct {
	Name     string    `json:"categoryName"`
	Roadmaps []Roadmap `json:"categoryRoadmaps"`
}
type Roadmap struct {
	Name         string       `json:"roadmapName"`
	Difficulties []Difficulty `json:"roadmapDifficulties"`
}

type Difficulty struct {
	Level  string  `json:"difficultyLevel"`
	Skills []Skill `json:"difficultySkills"`
}

type Skill struct {
	Title       string     `json:"skillTitle"`
	Description string     `json:"skillDescription"`
	Status      bool       `json:"skillStatus"`
	Resources   []Resource `json:"skillResources"`
}

type Resource struct {
	Name        string `json:"resourceName"`
	Description string `json:"resourceDescription"`
}
