package main

import (
	// system
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	// discord
	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/gateway"

	// .env lol
	"github.com/joho/godotenv"
)

func main() {
	slog.Info("Wiring")
	slog.Info("disgo version", slog.String("version", disgo.Version))

	err := godotenv.Load()

	if err != nil {
		slog.Error("Failed to wire :(")
		panic(err)
	}

	slog.Info("Wiring Complete.")

	client, err := disgo.New(os.Getenv("TOKEN"),
		// set gateway options
		bot.WithGatewayConfigOpts(
			// set enabled intents
			gateway.WithIntents(
				gateway.IntentGuilds,
				gateway.IntentGuildMessages,
				gateway.IntentDirectMessages,
			),
		),
		// add event listeners
		bot.WithEventListenerFunc(func(e *events.MessageCreate) {
			// event code here
		}),
	)
	if err != nil {
		panic(err)
	}
	// connect to the gateway
	if err = client.OpenGateway(context.TODO()); err != nil {
		panic(err)
	}

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	<-s
}
