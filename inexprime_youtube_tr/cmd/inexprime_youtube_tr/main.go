// File: cmd/inexprime_youtube_tr/main.go
package main

import (
	"encoding/json"
	"fmt"
    "os"
    "log"
	"youtube_tracker/pkg/youtube"
)

// Defining parser struct
type Video struct {
	Items []struct {
		Snippet struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"snippet"`
	} `json:"items"`
}

// Parsing video details
func parseVideoDetails(videoDetails string) (string, string, error) {
	var video Video
	err := json.Unmarshal([]byte(videoDetails), &video)
	if err != nil {
		return "", "", err
	}

	title := video.Items[0].Snippet.Title
	description := video.Items[0].Snippet.Description

	return title, description, nil
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file. Err: %s", err)
	}
	// TODO put inside secrets or somewhere there
	videoID := os.Getenv("VIDEO_ID")
	apiKey := os.Getenv("API_KEY")

	videoDetails, err := youtube.FetchVideoDetails(videoID, apiKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	title, description, err := parseVideoDetails(videoDetails)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Limit the description to the first 100 characters
	if len(description) > 200 {
		description = description[:200] + "..."
	}

	fmt.Println("Title:", title)
	fmt.Println("Description:", description)

}
