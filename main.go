package main

import (
	"fmt"
	harborModel "github.com/dukhyungkim/harbor-client/model"
	"github.com/gofiber/fiber/v2"
	"harago/cmd"
	"harago/config"
	"harago/gservice"
	"harago/gservice/gchat"
	"harago/handler"
	"log"
	"net/http"
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

	executor := cmd.NewExecutor()
	if err = executor.LoadCommands(cfg); err != nil {
		log.Fatalln(err)
	}

	gChat, err := gchat.NewGChat(gService, handler.NewDMHandler(executor), handler.NewRoomHandler(executor))
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

	app.Post("/message", func(ctx *fiber.Ctx) error {
		var event gchat.ChatEvent
		if err := ctx.BodyParser(&event); err != nil {
			log.Println(err)
		}
		return ctx.JSON(gChat.HandleMessage(&event))
	})

	app.Post("/harbor_notify", func(ctx *fiber.Ctx) error {
		var event harborModel.WebhookEvent
		if err := ctx.BodyParser(&event); err != nil {
			log.Println(err)
		}
		go handler.HandleHarborEvent(&event)
		return ctx.SendStatus(http.StatusOK)
	})

	return app
}
