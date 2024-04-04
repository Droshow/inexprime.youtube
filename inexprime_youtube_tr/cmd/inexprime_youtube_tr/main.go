// File: cmd/inexprime_youtube_tr/main.go
package main

import (
    "fmt"
    "youtube_tracker/pkg/youtube"
)

func main() {
    //TODO: has to go somewhere into by UI useable file
    videoID := "0hBQBhinxeU"
    //TODO: has to be hid somewhere in secrets
    apiKey := "AIzaSyCtebTX6Qn1m1DL6ZpsRBv-czqBW7PkpCk"

    videoDetails, err := youtube.FetchVideoDetails(videoID, apiKey)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(videoDetails)
}