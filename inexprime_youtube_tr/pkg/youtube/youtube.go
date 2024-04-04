// File: pkg/youtube/youtube.go
package youtube

import (
    "fmt"
    "io"
    "net/http"
)

func FetchVideoDetails(videoID string, apiKey string) (string, error) {
    url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/videos?id=%s&key=%s&part=snippet,contentDetails,statistics,status", videoID, apiKey)

    resp, err := http.Get(url)
    if err != nil {
        return "", fmt.Errorf("HTTP request failed with error: %w", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", fmt.Errorf("Failed to read response body with error: %w", err)
    }

    return string(body), nil
}