package main

import (
	"docgo/config"
	"docgo/gservice"
	"docgo/gservice/gchat"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
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

	gService, err := gservice.NewGService(opts.Credential)
	if err != nil {
		log.Fatalln(err)
	}

	gChat, err := gchat.NewGChat(gService)
	if err != nil {
		log.Fatalln(err)
	}

	app := setup(gChat)

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("Server startup ... %s\n", addr)
	log.Fatalln(app.Listen(addr))
}

func setup(gChat *gchat.GChat) *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Post("/", func(ctx *fiber.Ctx) error {
		var event gchat.ChatEvent
		if err := ctx.BodyParser(&event); err != nil {
			log.Println(err)
		}
		return ctx.JSON(gChat.HandleMessage(&event))
	})

	return app
}
