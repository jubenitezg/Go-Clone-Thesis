package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	. "readme_extraction/model"
)

var (
	inputPath *string
)

func init() {
	inputPath = flag.String("input", "", "Path to the input file")

}

func main() {
	flag.Parse()
	if *inputPath == "" {
		flag.PrintDefaults()
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
	var repositories []Repository
	err = json.Unmarshal(file, &repositories)
	if err != nil {
		fmt.Println(err)
		return
	}
	var repositoryDescriptions []RepositoryDescription
	for i := 0; i < len(repositories); i++ {
		fmt.Println(repositories[i])
		//cmd := exec.Command("bash", "-c", fmt.Sprintf("git remote show %s | sed -n '/HEAD branch/s/.*: //p' | tr -d '\\n'", repositories[i].URL))
		//branch, err := cmd.Output()
		//if err != nil {
		//	fmt.Println(err)
		//	return
		//}
		//resp, err := http.Get("https://raw.githubusercontent.com/" + repositories[i].Owner + "/" + repositories[i].Name + "/" + string(branch) + "/README.md")
		//if err != nil {
		//	fmt.Println(err)
		//	return
		//}
		//fmt.Println(resp.StatusCode)
		//all, _ := io.ReadAll(resp.Body)
		//re := regexp.MustCompile(`\r?\n`)
		//descs := re.ReplaceAllString(string(all), " ")
		//fullDescription := repositories[i].Description + descs + "TOPICS:" + fmt.Sprint(repositories[i].Topics)
		//readme := string(all)
		repositoryDescriptions = append(repositoryDescriptions, RepositoryDescription{
			Name:        repositories[i].Name,
			Owner:       repositories[i].Owner,
			Description: repositories[i].Description,
			Topics:      repositories[i].Topics,
			//Readme:      descs,
			//ReadmeBytes: all,
		})
		//defer resp.Body.Close()
	}
	csvFile, err := os.Create("output.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer csvFile.Close()
	csvwriter := csv.NewWriter(csvFile)
	csvwriter.Write([]string{"name", "owner", "description", "topics"})
	defer csvwriter.Flush()
	// removed readme from the csv since there are some issues with length of the readme
	for _, repositoryDescription := range repositoryDescriptions {
		//var b bytes.Buffer
		//gz := gzip.NewWriter(&b)
		//if _, err := gz.Write(repositoryDescription.ReadmeBytes); err != nil {
		//	log.Fatal(err)
		//}
		//if err := gz.Close(); err != nil {
		//	log.Fatal(err)
		//}
		//fmt.Println(b.Bytes())
		err := csvwriter.Write([]string{repositoryDescription.Name, repositoryDescription.Owner, repositoryDescription.Description, fmt.Sprint(repositoryDescription.Topics)})
		if err != nil {
			panic(err)
		}
		fmt.Println(repositoryDescription)
	}
	fmt.Println("Done")
}
