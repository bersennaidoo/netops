package models

type Configuration struct {
	Hostname []string            `json:"hostname"`
	Username []string            `json:"username"`
	Password []string            `json:"password"`
	Configs  map[string][]string `json:"configs"`
}
