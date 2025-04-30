package models

type Blueprint struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Steps       []Step `yaml:"steps"`
}

// Step = Sequential unit of executable tasks
type Step struct {
	Label string `yaml:"label"`
	Tasks []Task `yaml:"tasks"`
}

// Task = Single executable action
type Task struct {
	Label string                 `yaml:"label"`
	Uses  string                 `yaml:"uses"`
	With  map[string]interface{} `yaml:"with"`
}
