package common

type CodeFragment struct {
	Code string `json:"code"`
	Id   string `json:"id"`
	Line int    `json:"line"`
	Path string `json:"path"`
}
