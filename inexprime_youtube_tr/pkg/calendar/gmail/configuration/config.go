package configuration

import (
    "log"
    "os"

    "github.com/joho/godotenv"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
)

type Config struct {
    GoogleLoginConfig oauth2.Config
}

var AppConfig Config

func LoadConfig() Config {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file. Err: %s", err)
    }

    AppConfig.GoogleLoginConfig = oauth2.Config{
        RedirectURL:  os.Getenv("REDIRECT_URL"),
        ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
        ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
        Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
            "https://www.googleapis.com/auth/userinfo.profile"},
        Endpoint: google.Endpoint,
    }

    return AppConfig
}