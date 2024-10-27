package api

import (
	"encoding/json"
	"fmt"
)

func AcceptChallenge(id string) error {
	return Post(fmt.Sprintf("https://lichess.org/api/challenge/%s/accept", id), token)
}

func DeclineChallenge(id string) error {
	return Post(fmt.Sprintf("https://lichess.org/api/challenge/%s/decline", id), token)
}

func MakeMove(id string, move string) error {
	return Post(fmt.Sprintf("https://lichess.org/api/bot/game/%s/move/%s", id, move), token)
}

type Event struct {
	Type    string
	Content string
}

func StreamEvents() (chan Event, error) {
	stream, err := Stream("https://lichess.org/api/stream/event", token)
	if err != nil {
		return nil, err
	}
	ch := make(chan Event)
	go func() {
		for line := range stream {
			var event Event
			json.Unmarshal([]byte(line), &event)
			event.Content = line
			ch <- event
		}
	}()
	return ch, nil
}

func StreamGame(id string) (chan Event, error) {
	stream, err := Stream(fmt.Sprintf("https://lichess.org/api/bot/game/stream/%s", id), token)
	if err != nil {
		return nil, err
	}
	ch := make(chan Event)
	go func() {
		for line := range stream {
			var event Event
			json.Unmarshal([]byte(line), &event)
			event.Content = line
			ch <- event
		}
	}()
	return ch, nil
}

type ProfileInfo struct {
	ID       string
	Username string
}

func GetProfile() (ProfileInfo, error) {
	res, err := Get("https://lichess.org/api/account", token)
	if err != nil {
		return ProfileInfo{}, err
	}
	var profile ProfileInfo
	json.Unmarshal([]byte(res), &profile)
	return profile, nil
}
