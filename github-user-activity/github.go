package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const GitHubAPIUrl = "https://api.github.com"

func fetchEvent(username string, page int) ([]GitHubEvent, error) {
	const EventsPerPage = 5

	url := fmt.Sprintf("%s/users/%s/events?page=%d&per_page=%d", GitHubAPIUrl, username, page, EventsPerPage)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	client := &http.Client{Timeout: 10 * time.Second}
	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %s", response.Status)
	}

	var events []GitHubEvent

	if err := json.NewDecoder(response.Body).Decode(&events); err != nil {
		return nil, err
	}

	return events, nil
}

func printPrettyJSON(data any) {
	// MarshalIndent(data, prefix, indent)
	// We leave prefix empty and use 4 spaces for indent
	prettyJSON, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		fmt.Printf("Error formatting JSON: %v\n", err)
		return
	}

	fmt.Printf("\033[32m%s\033[0m\n", string(prettyJSON)) // Prints everything in Green
}
