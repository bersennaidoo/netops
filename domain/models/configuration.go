package models

type Configuration struct {
	Hostname string   `json:"hostname"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Config   []string `json:"config"`
}
