package main

type GitHubEvent struct {
	Type      string `json:"type"`
	CreatedAt string `json:"created_at"`
	Actor     struct {
		Login        string `json:"login"`
		DisplayLogin string `json:"display_login"`
		ID           int    `json:"id"`
		URL          string `json:"url"`
	} `json:"actor"`
	Repo struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"repo"`
}
