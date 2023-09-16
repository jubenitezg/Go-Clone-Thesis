package main

import (
	"fmt"
	"go-extractor/extractor"
)

func main() {
	filename := "test.go"
	ex, err := extractor.NewExtractor(filename)
	if err != nil {
		fmt.Println("Error creating extractor:", err)
		return
	}
	myPaths := ex.GeneratePathForFunctions()
	fmt.Println("=========================================")
	compPaths := ex.GeneratePathForFunctionsCompare()
	if len(myPaths) != len(compPaths) {
		fmt.Println("Different number of paths")
		return
	}
	for i := 0; i < len(myPaths); i++ {
		if myPaths[i] != compPaths[i] {
			fmt.Println("Different paths")
			return
		}
	}
	fmt.Println("Same paths")

}
