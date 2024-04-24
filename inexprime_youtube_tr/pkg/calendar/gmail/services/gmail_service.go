package services

import (
	"context"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"youtube_tracker/pkg/calendar/gmail/configuration"
)

type GmailService struct {
	Service *calendar.Service
}

func NewGmailService(ctx context.Context, token *oauth2.Token) (*GmailService, error) {
	client := configuration.AppConfig.GoogleLoginConfig.Client(ctx, token)
	service, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}
	return &GmailService{Service: service}, nil
}

func (s *GmailService) CreateEvent(ctx context.Context, event *calendar.Event) (*calendar.Event, error) {
	newEvent, err := s.Service.Events.Insert("primary", event).Do()
	if err != nil {
		return nil, err
	}
	return newEvent, nil
}
