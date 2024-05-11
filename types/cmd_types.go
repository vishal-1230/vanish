package vanish_types

import "time"

// Any is a type that can hold any value
type Any interface{}

// vanish config file
type VanishConfig struct {
	Models          []Model          `json:"models"`
	ServicesRunning []ServiceRunning `json:"services"`
}

type Model struct {
	Name         string     `json:"name"`
	HFURL        string     `json:"hf-url"`
	Transformer  bool       `json:"transformer"`
	Tokenizer    bool       `json:"tokenizer"`
	Architecture string     `json:"architecture"`
	Task         string     `json:"task"`
	Framework    string     `json:"framework"`
	Language     string     `json:"language"`
	License      string     `json:"license"`
	Description  string     `json:"description"`
	Parameters   Parameters `json:"parameters"`
	DownloadedOn string     `json:"downloaded_on"`
	Size         string     `json:"size"`
}

type Parameters struct {
	NumLabels int `json:"num_labels"`
}

type ServiceRunning struct {
	ModelId   string    `json:"model-id"`
	Name      string    `json:"model-name"`
	Port      int       `json:"port"`
	StartedAt time.Time `json:"started-at"`
}
