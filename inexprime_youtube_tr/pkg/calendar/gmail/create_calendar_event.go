package gmail

import (
	"context"
	"encoding/json"
	"os"
	"youtube_tracker/pkg/calendar/gmail/services"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
)

func RegisterEventRoutes(app *fiber.App) {
	app.Post("/create_event", CreateEvent)
}

func CreateEvent(c *fiber.Ctx) error {
	token, err := getTokenFromStorage()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve token: " + err.Error())
	}

	gs, err := services.NewGmailService(context.Background(), token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to initialize Gmail service: " + err.Error())
	}

	var event calendar.Event
	if err := c.BodyParser(&event); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error parsing event data: " + err.Error())
	}

	createdEvent, err := gs.CreateEvent(context.Background(), &event)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create event: " + err.Error())
	}

	return c.JSON(createdEvent)
}

func getTokenFromStorage() (*oauth2.Token, error) {
	file, err := os.Open("token.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	token := &oauth2.Token{}
	err = json.NewDecoder(file).Decode(token)
	if err != nil {
		return nil, err
	}

	return token, nil
}
