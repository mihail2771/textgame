package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Game struct {
	Rooms   map[RoomID]*Room
	Player  Player
	Aliases map[string]string
}

func loadGame(path string) (*Game, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	type playerConfig struct {
		Room  RoomID `json:"room"`
		Items []Item `json:"items"`
		Isbag bool   `json:"isbag"`
	}

	type gameConfig struct {
		Rooms   map[RoomID]*Room  `json:"rooms"`
		Player  playerConfig      `json:"player"`
		Aliases map[string]string `json:"aliases"`
	}

	var config gameConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	for id, room := range config.Rooms {
		room.ID = id
	}

	playerRoom, ok := config.Rooms[config.Player.Room]
	if !ok {
		return nil, fmt.Errorf("unknown player room %q", config.Player.Room)
	}

	return &Game{
		Rooms: config.Rooms,
		Player: Player{
			Room:  playerRoom,
			Items: config.Player.Items,
			Isbag: config.Player.Isbag,
		},
		Aliases: config.Aliases,
	}, nil
}

func SubEmpty[T comparable](value, defaultValue T) T {
	var zero T
	if value == zero {
		return defaultValue
	}
	return value
}

func parseItem(label string) (Item, bool) {
	item, ok := itemLabelToID[label]
	return item, ok
}

func parseRoom(label string) (RoomID, bool) {
	roomID, ok := roomLabelToID[label]
	return roomID, ok
}

var game *Game

func initGame() {
	config, err := loadGame("game_init.json")
	if err != nil {
		panic(err)
	}

	game = config
}

func updateWorld() {
	if game.Player.Room.ID == RoomStreet {
		return
	}
	if game.Player.Isbag {
		game.Rooms[RoomKitchen].TaskDescription = taskGoUniversity
	}
}

func handleCommand(command string) string {
	args := strings.Fields(command)
	if len(args) == 0 {
		return msgUnknownCommand
	}

	var result string
	switch args[0] {
	case cmdLook:
		result = game.Player.lookAround()
	case cmdGo:
		if len(args) < 2 {
			return msgNoPath
		}
		roomID, ok := parseRoom(args[1])
		if !ok {
			return msgNoSuch
		}
		result = game.Player.goTo(roomID)
	case cmdWear:
		if len(args) < 2 {
			return msgNoWear
		}
		itemID, ok := parseItem(args[1])
		if !ok || itemID != ItemBackpack {
			return msgNoSuch
		}
		result = game.Player.getBag()
	case cmdTake:
		if len(args) < 2 {
			return msgNoTake
		}
		itemID, ok := parseItem(args[1])
		if !ok {
			return msgNoSuch
		}
		if game.Player.Isbag {
			result = game.Player.getInBag(itemID)
		} else {
			result = msgNoBagSpace
		}
	case cmdApply:
		if len(args) < 3 {
			return msgNoApply
		}
		itemID, ok := parseItem(args[1])
		if !ok {
			result = msgNoItemInInventory(args[1])
		} else {
			result = game.Player.activeItem(itemID, args[2])
		}
	default:
		result = msgUnknownCommand
	}

	updateWorld()
	return result
}
