package model

type RepositoryDescription struct {
	Name        string   `csv:"name"`
	Owner       string   `csv:"owner"`
	Description string   `csv:"description"`
	Readme      string   `csv:"readme"`
	Topics      []string `csv:"topics"`
	ReadmeBytes []byte
}
