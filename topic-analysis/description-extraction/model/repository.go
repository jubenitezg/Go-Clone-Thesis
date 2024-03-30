package model

type Repository struct {
	Name        string   `json:"name"`
	Owner       string   `json:"owner"`
	Description string   `json:"description"`
	Topics      []string `json:"topics"`
	URL         string   `json:"url"`
}
