package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: github-activity <username>")
		return
	}

	username := os.Args[1]
	page := 1

	if len(os.Args) > 2 {
		var err error
		page, err = strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println("Defaulting to 1st page.")
			page = 1
		}
	}

	events, err := fetchEvent(username, page)

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("Fetching activity for %s (Page %d)...\n", username, page)

	printPrettyJSON(events)
}
