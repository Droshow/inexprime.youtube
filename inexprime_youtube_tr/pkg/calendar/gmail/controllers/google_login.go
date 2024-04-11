package controllers

import (
    "github.com/Droshow/inexprime.youtube/inexprime_youtube_tr/pkg/calendar/gmail/configuration"
    "github.com/gofiber/fiber/v2"
)

func GoogleLogin(c *fiber.Ctx) error {
    url := configuration.AppConfig.GoogleLoginConfig.AuthCodeURL("randomstate")

    c.Status(fiber.StatusSeeOther)
    c.Redirect(url)
    return c.JSON(url)
}