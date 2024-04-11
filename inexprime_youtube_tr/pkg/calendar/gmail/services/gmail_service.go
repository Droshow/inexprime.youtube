package services

import (
	"context"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"github.com/Droshow/inexprime.youtube/inexprime_youtube_tr/pkg/calendar/gmail/configuration"
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