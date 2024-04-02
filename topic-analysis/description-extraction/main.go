package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	. "readme_extraction/model"
)

var (
	inputPath  *string
	outputDir  *string
	singleLine *bool
)

func init() {
	inputPath = flag.String("input-path", "", "Path to the input file")
	outputDir = flag.String("output-directory", "", "Path to the output directory")
	singleLine = flag.Bool("single-line", true, "Whether to output the readme in a single line (default)")
}

func main() {
	flag.Parse()
	if *inputPath == "" || *outputDir == "" {
		flag.Usage()
		return
	}
	if _, err := os.Stat(*inputPath); os.IsNotExist(err) {
		fmt.Printf("File %s does not exist\n", *inputPath)
		return
	}
	file, err := os.ReadFile(*inputPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	var repositories []*Repository
	err = json.Unmarshal(file, &repositories)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, repository := range repositories[:2] {
		cmd := exec.Command("bash", "-c", fmt.Sprintf("git remote show %s | sed -n '/HEAD branch/s/.*: //p' | tr -d '\\n'", repository.URL))
		branch, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
			return
		}
		resp, err := http.Get("https://raw.githubusercontent.com/" + repository.Owner + "/" + repository.Name + "/" + string(branch) + "/README.md")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(resp.StatusCode)
		readmeBytes, _ := io.ReadAll(resp.Body)
		defer resp.Body.Close()
		repository.ReadmeBase64 = base64.StdEncoding.EncodeToString(readmeBytes)
	}

	buffer := new(bytes.Buffer)
	enc := json.NewEncoder(buffer)
	enc.SetEscapeHTML(false)
	if !*singleLine {
		enc.SetIndent("", "  ")
	}
	err = enc.Encode(repositories)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	err = os.WriteFile(*outputDir+"/output.json", buffer.Bytes(), 0644)
}
