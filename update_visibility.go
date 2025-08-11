package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

const githubAPI = "https://api.github.com/repos/"

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run update_visibility.go <visibility> <owner/repo> [owner/repo] ...")
		os.Exit(1)
	}

	visibility := os.Args[1]
	repos := os.Args[2:]

	err := loadEnv(".env")
	if err != nil {
		log.Fatalf("Failed to load .env: %v", err)
	}

	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		fmt.Println("Error: GITHUB_TOKEN environment variable not set")
		os.Exit(1)
	}

	for _, repo := range repos {
		err := updateRepoVisibility(repo, visibility, token)
		if err != nil {
			fmt.Printf("Failed to update %s: %v\n", repo, err)
		} else {
			fmt.Printf("Updated %s to %s\n", repo, visibility)
		}
	}
}

func updateRepoVisibility(repo, visibility, token string) error {
	url := githubAPI + repo

	payload := map[string]string{"visibility": visibility}
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/vnd.github.nebula-preview+json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return fmt.Errorf("GitHub API returned status %s", resp.Status)
	}

	return nil
}

func loadEnv(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Split key=value
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		os.Setenv(key, value)
	}

	return scanner.Err()
}
