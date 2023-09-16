package main

import (
	"fmt"
	"go-extractor/extractor"
)

func main() {
	filename := "extractor/test_assets/test.go"
	ex, err := extractor.NewExtractor(filename, false)
	if err != nil {
		fmt.Println("Error creating extractor:", err)
		return
	}
	//ex.GenerateProgramAstPaths()
	//fmt.Println("=========================================")
	ex.GeneratePathForFunctionsCompare()
	//if len(myPaths) != len(compPaths) {
	//	fmt.Println("Different number of paths")
	//	return
	//}
	//for i := 0; i < len(myPaths); i++ {
	//	if myPaths[i] != compPaths[i] {
	//		fmt.Println("Different paths")
	//		return
	//	}
	//}
	//fmt.Println("Same paths")

}
