package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/Droshow/inexprime.youtube/inexprime_youtube_tr/pkg/calendar/gmail/controllers"
)

func main() {
    app := fiber.New()

    app.Get("/google_login", controllers.GoogleLogin)
    app.Get("/google_callback", controllers.GoogleCallback)

    app.Listen(":8080")
}