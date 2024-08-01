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
	"github.com/disgoorg/disgo/discord"
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
		slog.Error("Failed to wire :(")
		panic(err)
	}
	// connect to the gateway
	if err = client.OpenGateway(context.TODO()); err != nil {
		slog.Error("Failed to wire :(")
		panic(err)
	}

	slog.Info("Wiring Complete.")
	slog.Info("CTRL-C to stop wiring")

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	<-s
}

func onMessageCreate(event *events.MessageCreate) {
	if event.Message.Author.Bot {
		return
	}

	var message string

	if event.Message.Content == "ping" {
		message = "pong"
	} else if event.Message.Content == "pong" {
		message = "are you retarded?"
	}

	if message != "" {
		_, _ = event.Client().Rest().CreateMessage(event.ChannelID, discord.NewMessageCreateBuilder().SetContent(message).Build())
	}
}
