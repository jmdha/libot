package main

import (
	"encoding/json"
	"fmt"
	"github.com/phsym/console-slog"
	"libot/api"
	"log/slog"
	"os"
)

func InitLog() {
	logger := slog.New(
		console.NewHandler(os.Stdout, &console.HandlerOptions{AddSource: true, Level: slog.LevelDebug}),
	)
	slog.SetDefault(logger)
}

func main() {
	InitLog()
	token := os.Args[1]
	path := os.Args[2]
	slog.Info("Retrieveing profile")
	profile, err := api.GetProfile(token)
	if err != nil {
		slog.Error("Failed to retrieve profile")
		panic(err)
	}
	slog.Info("Attempting to connect to event stream")
	stream, err := api.StreamEvents(token)
	if err != nil {
		slog.Error("Failed to initialize event stream")
		panic(err)
	}
	slog.Info("Event stream successfully connected")
	for event := range stream {
		content := event.Content
		switch event.Type {
		case "gameStart":
			var event api.GameStartEvent
			json.Unmarshal([]byte(content), &event)
			slog.Info("Starting game " + event.Game.ID)
			go HandlGame(token, event.Game.ID, profile.ID, path)
		case "gameFinish":
			var event api.GameFinsihEvent
			json.Unmarshal([]byte(content), &event)
			slog.Info("Game finished " + event.Game.ID)
		case "challenge":
			var event api.ChallengeEvent
			json.Unmarshal([]byte(content), &event)
			slog.Info(fmt.Sprintf("Received challenge from %s (%d)", event.Challenge.Challenger.Name, event.Challenge.Challenger.Rating))
			slog.Info("Accepting challenge")
			api.AcceptChallenge(event.Challenge.ID, token)
		case "challengeCanceled":
			var event api.ChallengeCanceledEvent
			json.Unmarshal([]byte(content), &event)
		case "challengeDeclined":
			var event api.ChallengeDeclinedEvent
			json.Unmarshal([]byte(content), &event)
		}
	}
}


