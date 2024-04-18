package main

import (
	"fmt"
	"metadata-extraction/extractor"
	"metadata-extraction/model"
)

// URL:      "https://github.com/go-xorm/xorm", disabled repo
func main() {
	testRepo := &model.Repository{
		Name:      "hub",
		Owner:     "mislav",
		URL:       "https://github.com/mislav/hub",
		Issues:    2018,
		Stars:     22649,
		License:   "MIT",
		CreatedAt: "2009-12-05",
	}
	//testRepo := &model.Repository{
	//	Name:      "xorm",
	//	Owner:     "go-xorm",
	//	URL:       "https://github.com/go-xorm/xorm",
	//	Issues:    2018,
	//	Stars:     22649,
	//	License:   "MIT",
	//	CreatedAt: "2009-12-05",
	//}
	metadata, err := extractor.GetRepositoryMetadata(testRepo)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(metadata)
}
