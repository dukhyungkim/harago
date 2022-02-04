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
	"harago/repo"
	"harago/stream"
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	log.Println("Harago Starting.")

	opts, err := config.ParseFlags()
	if err != nil {
		log.Fatalln(err)
	}

	cfg, err := config.NewConfig(opts)
	if err != nil {
		log.Fatalln(err)
	}

	db, err := repo.NewPostgres(cfg.DB)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("connect to postgres ... success")

	if opts.AutoMigration {
		if err := db.AutoMigration(); err != nil {
			log.Fatalln(err)
		}
	}
	log.Println("migrate to postgres ... success")

	etcd, err := repo.NewEtcd(cfg.Etcd)
	if err != nil {
		log.Fatalln(err)
	}
	defer etcd.Close()
	log.Println("connect to etcd ... success")

	gService, err := gservice.NewGService(opts.Credential)
	if err != nil {
		log.Fatalln(err)
	}

	executor := cmd.NewExecutor()
	if err = executor.LoadCommands(cfg); err != nil {
		log.Fatalln(err)
	}

	dmHandler := handler.NewDMHandler(executor, db)
	roomHandler := handler.NewRoomHandler(executor, db)
	gChat, err := gchat.NewGChat(gService, dmHandler, roomHandler)
	if err != nil {
		log.Fatalln(err)
	}

	streamClient, err := stream.NewStreamClient(cfg.Nats)
	if err != nil {
		log.Fatalln(err)
	}
	defer streamClient.Close()
	log.Println("connect to nats ... success")

	respHandler := handler.NewResponseHandler(gChat, db)
	if err = streamClient.ClamResponse(respHandler.NotifyResponse); err != nil {
		log.Fatalln(err)
	}

	harborEventHandle := handler.NewHarborEventHandler(streamClient)

	app := setup(gChat, harborEventHandle)

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("Server startup ... %s\n", addr)
	log.Fatalln(app.Listen(addr))
}

func setup(gChat *gchat.GChat, harborEventHandler *handler.HarborEventHandler) *fiber.App {
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
		go harborEventHandler.HandleHarborEvent(&event)
		return ctx.SendStatus(http.StatusOK)
	})

	return app
}
