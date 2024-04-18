package model

type Metadata struct {
	Owner        string         `json:"owner"`
	Name         string         `json:"name"`
	URL          string         `json:"url"`
	Contributors int            `json:"contributors"`
	Commits      int            `json:"commits"`
	CreatedAt    string         `json:"createdAt"`
	Stars        int            `json:"stars"`
	Languages    map[string]int `json:"languages"`
	License      string         `json:"license"`
	Issues       int            `json:"issues"`
	// LOC
}

type Repository struct {
	Name      string `json:"name"`
	Owner     string `json:"owner"`
	URL       string `json:"url"`
	Issues    int    `json:"issues"`
	CreatedAt string `json:"createdAt"`
	Stars     int    `json:"stars"`
	License   string `json:"license"`
}
