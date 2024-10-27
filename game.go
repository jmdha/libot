package main

import (
	"encoding/json"
	"libot/api"
	"log/slog"
	"strings"
)

func HandlGame(id string, bot string, path string) {
	engine := NewEngine(path)
	engine.Init()
	stream, err := api.StreamGame(id)
	if err != nil {
		slog.Error("Failed to stream game " + id)
		panic(err)
	}
	var initialFEN string
	var isWhite bool
	for event := range stream {
		content := event.Content
		switch event.Type {
		case "gameFull":
			var event api.GameFullEvent
			json.Unmarshal([]byte(content), &event)
			initialFEN = event.InitialFen
			isWhite = bot == event.White.ID
			HandleTurn(&engine, id, isWhite, initialFEN, event.State)
		case "gameState":
			var event api.GameStateEvent
			json.Unmarshal([]byte(content), &event)
			if event.Status != "started" {
				slog.Info("Game " + id + " terminated with status " + event.Status)
				return
			}
			HandleTurn(&engine, id, isWhite, initialFEN, event)
		case "chatLine":
			var event api.ChatLineEvent
			json.Unmarshal([]byte(content), &event)
		case "opponentGone":
			var event api.OpponentGoneEvent
			json.Unmarshal([]byte(content), &event)
		}
	}
}

func HandleTurn(engine *Engine, gameID string, isWhite bool, initialFen string, gameState api.GameStateEvent) error {
	if IsTurn(initialFen, gameState.Moves, isWhite) {
		move := engine.BestMove(initialFen, gameState.Moves, gameState.WTime, gameState.BTime)
		api.MakeMove(gameID, move)
	}
	return nil
}

func IsTurn(fen string, moves string, isWhite bool) bool {
	startingWhite := IsWhiteStarting(fen)
	if len(moves) == 0 {
		return isWhite == startingWhite
	}
	offset := len(strings.Split(moves, " "))
	currentWhite := startingWhite && (offset%2 == 0) ||
		!startingWhite && (offset%2 == 1)
	return isWhite == currentWhite
}

func IsWhiteStarting(fen string) bool {
	if fen == "startpos" {
		return true
	}
	return strings.Contains(fen, "w")
}
