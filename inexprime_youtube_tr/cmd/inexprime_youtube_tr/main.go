package main

import (
    "github.com/joho/godotenv"
    "log"
    "os"
    "youtube_tracker/pkg/youtube"
    "youtube_tracker/pkg/calendar/gmail"
    "google.golang.org/api/calendar/v3"
)

func main() {
    err := godotenv.Load("../../.env")
    if err != nil {
        log.Fatalf("Error loading .env file. Err: %s", err)
    }

    videoID := "0hBQBhinxeU"
    apiKey := os.Getenv("API_KEY")

    title, description, err := youtube.FetchVideoDetails(videoID, apiKey)
    if err != nil {
        log.Println(err)
        return
    }
    // Limit the description to the first 200 characters
    if len(description) > 200 {
        description = description[:200] + "..."
    }
    event := &calendar.Event{
        Summary: title,
        Description: description,
        // ... (set other fields as needed) ...
    }
    err = gmail.CreateEventWithoutCtx(event)
    if err != nil {
        log.Println(err)
        return
    }
    log.Println("Event created successfully")
}