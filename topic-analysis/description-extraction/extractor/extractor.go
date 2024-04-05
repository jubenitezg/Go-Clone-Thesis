package extractor

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"readme_extraction/model"
	"strconv"
	"strings"
)

type ReadmeExtractor struct {
	InputPath *string
	OutputDir *string
	From      *int
	To        *int
}

func NewReadmeExtractor(inputPath, outputDir *string, from, to *int) *ReadmeExtractor {
	return &ReadmeExtractor{
		InputPath: inputPath,
		OutputDir: outputDir,
		From:      from,
		To:        to,
	}
}

func (re *ReadmeExtractor) Extract() ([]model.Repository, error) {
	if _, err := os.Stat(*re.InputPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("file %s does not exist", *re.InputPath)
	}
	file, err := os.ReadFile(*re.InputPath)
	if err != nil {
		return nil, err
	}
	var repositories []*model.Repository
	err = json.Unmarshal(file, &repositories)
	if err != nil {
		return nil, err
	}
	if *re.To == -1 {
		*re.To = len(repositories)
	}
	repositoriesReadme := make([]model.Repository, 0)
	for i, repository := range repositories[*re.From:*re.To] {
		fmt.Println("Repository #" + fmt.Sprint(*re.From+i) + ": " + repository.Owner + "/" + repository.Name)
		readme64, err := getReadmeBase64(repository)
		if err != nil {
			return repositoriesReadme, err
		}
		repositoriesReadme = append(repositoriesReadme, model.Repository{
			Name:         repository.Name,
			Owner:        repository.Owner,
			Description:  repository.Description,
			Topics:       repository.Topics,
			URL:          repository.URL,
			ReadmeBase64: readme64,
		})
	}
	return repositoriesReadme, nil
}

func getReadmeBase64(repository *model.Repository) (string, error) {
	remaining := verifyRateLimit()
	fmt.Println("Remaining rate limit: " + fmt.Sprint(remaining))
	if remaining <= 0 {
		return "", fmt.Errorf("rate limit exceeded")
	}
	cmd := exec.Command("gh", "api", "-H", "Accept: application/vnd.github+json",
		"-H", "X-GitHub-Api-Version: 2022-11-28", fmt.Sprintf("/repos/%s/%s/readme", repository.Owner, repository.Name),
		"-q", ".content")
	outputBytes, err := cmd.Output()
	output := string(outputBytes)
	if err != nil {
		// repo has no readme or repo is disabled
		if strings.Contains(output, "Not Found") || strings.Contains(output, "access blocked") {
			return "", nil
		}
		return "", err
	}
	return output, nil
}

func verifyRateLimit() int {
	cmd := exec.Command("gh", "api", "rate_limit", "-q", ".resources.core.remaining")
	output, err := cmd.Output()
	if err != nil {
		return -1
	}
	rawOutput := strings.TrimSuffix(string(output), "\n")
	atoi, err := strconv.Atoi(rawOutput)
	if err != nil {
		return -1
	}
	return atoi
}
