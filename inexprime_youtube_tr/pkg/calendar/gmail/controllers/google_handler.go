package controllers

import (
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"youtube_tracker/pkg/calendar/gmail/configuration"
)

func GoogleLogin(c *fiber.Ctx) error {
	state := uuid.NewString() // Generate a new UUID for the state
	c.Cookie(&fiber.Cookie{   // Store state in a secure, HTTP-only cookie
		Name:     "oauthstate",
		Value:    state,
		HTTPOnly: true,
		Secure:   true, // Set to true if using HTTPS
	})
	url := configuration.AppConfig.GoogleLoginConfig.AuthCodeURL(state)

	c.Status(fiber.StatusSeeOther)
	c.Redirect(url)
	return nil
}

func GoogleCallback(c *fiber.Ctx) error {
	// Retrieve the state from the user's cookies to compare with the state in the query.
	state := c.Cookies("oauthstate")
	queryState := c.Query("state")

	// Check if the state from the cookie matches the state returned in the query.
	// This is a security measure to prevent CSRF attacks.
	if state != queryState {
		return c.SendString("States don't match!")
	}

	// Retrieve the authorization code from the query parameters.
	code := c.Query("code")

	// Exchange the authorization code for an access token.
	token, err := configuration.AppConfig.GoogleLoginConfig.Exchange(context.Background(), code)
	if err != nil {
		return c.SendString("Code-token exchange failed: " + err.Error())
	}

	// Create a new file to store the token for future use.
	file, err := os.Create("token.json")
	if err != nil {
		return c.SendString("Failed to create token file: " + err.Error())
	}
	// Ensure the file is closed after the function exits.
	defer file.Close()

	// Write the token to the file in JSON format.
	if err := json.NewEncoder(file).Encode(token); err != nil {
		return c.SendString("Failed to write token to file: " + err.Error())
	}

	// Make a request to Google's userinfo endpoint to fetch user data using the access token.
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return c.SendString("User data fetch failed: " + err.Error())
	}
	// Ensure the HTTP response body is closed after the function exits.
	defer resp.Body.Close()

	// Read the user data from the response body.
	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.SendString("JSON parsing failed: " + err.Error())
	}

	// Return the user data as a string.
	return c.SendString(string(userData))
}
