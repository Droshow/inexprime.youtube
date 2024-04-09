// File: cmd/inexprime_youtube_tr/main.go
package main

import (
	"encoding/json"
	"fmt"
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
	//TODO: has to go somewhere into by UI useable file
	videoID := "0hBQBhinxeU"
	//TODO: has to be hidden somewhere in secrets
	apiKey := "AIzaSyCtebTX6Qn1m1DL6ZpsRBv-czqBW7PkpCk"

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
