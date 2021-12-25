package main

import (
	"context"
	"docgo/config"
	"fmt"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/chat/v1"
	"google.golang.org/api/option"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	log.Println("Docgo Starting.")

	opts, err := config.ParseFlags()
	if err != nil {
		log.Fatalln(err)
	}

	cfg, err := config.NewConfig(opts)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(cfg)

	googleClient, err := NewGoogleClient(opts.Credential)
	if err != nil {
		log.Fatalln(err)
	}

	chatService, err := NewChatService(googleClient)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(*chatService)
}

func NewGoogleClient(credential string) (*http.Client, error) {
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

	return httpClient, nil
}

func NewChatService(client *http.Client) (*chat.Service, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return chat.NewService(ctx, option.WithHTTPClient(client))
}
