package controllers

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/Droshow/inexprime.youtube/inexprime_youtube_tr/pkg/calendar/gmail/configuration"
	"github.com/gofiber/fiber/v2"
)

func GoogleCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	if state != "randomstate" {
		return c.SendString("States don't Match!!")
	}

	code := c.Query("code")

	token, err := configuration.AppConfig.GoogleLoginConfig.Exchange(context.Background(), code)
	if err != nil {
		return c.SendString("Code-Token Exchange Failed")
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return c.SendString("User Data Fetch Failed")
	}

	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.SendString("JSON Parsing Failed")
	}

	return c.SendString(string(userData))
}