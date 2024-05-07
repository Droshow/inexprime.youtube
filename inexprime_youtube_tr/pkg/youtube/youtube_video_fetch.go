// File: pkg/youtube/youtube_video_fetch.go
package youtube

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

type Video struct {
    Items []struct {
        Snippet struct {
            Title       string `json:"title"`
            Description string `json:"description"`
        } `json:"snippet"`
    } `json:"items"`
}

func parseVideoDetails(videoDetails string) (string, string, error) {
    var video Video
    err := json.Unmarshal([]byte(videoDetails), &video)
    if err != nil {
        return "", "", err
    }

    if len(video.Items) == 0 {
        return "", "", fmt.Errorf("no items in video")
    }

    title := video.Items[0].Snippet.Title
    description := video.Items[0].Snippet.Description

    return title, description, nil
}

func FetchVideoDetails(videoID string, apiKey string) (string, string, error) {
    url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/videos?id=%s&key=%s&part=snippet,contentDetails,statistics,status", videoID, apiKey)

    resp, err := http.Get(url)
    if err != nil {
        return "", "", fmt.Errorf("HTTP request failed with error: %w", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", "", fmt.Errorf("Failed to read response body with error: %w", err)
    }

    title, description, err := parseVideoDetails(string(body))
    if err != nil {
        return "", "", err
    }

    return title, description, nil
}