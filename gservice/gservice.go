package gservice

import (
	"context"
	"fmt"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"
	"time"
)

type GService struct {
	client *http.Client
}

func NewGService(credential string) (*GService, error) {
	b, err := ioutil.ReadFile(credential)
	if err != nil {
		return nil, fmt.Errorf("unable to read client secret file: %w", err)
	}

	const ChatScope = "https://www.googleapis.com/auth/chat.bot"
	var scope = []string{ChatScope}
	configFromJSON, err := google.JWTConfigFromJSON(b, scope...)
	if err != nil {
		return nil, fmt.Errorf("unable to parse client secret file to configFromJSON: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	httpClient := configFromJSON.Client(ctx)

	return &GService{client: httpClient}, nil
}

func (svc *GService) GetClient() *http.Client {
	return svc.client
}
