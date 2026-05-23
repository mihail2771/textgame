package main

import (
	"strings"
)

var positions = []string{"на столе", "на стуле"}

var rooms map[string]*Room
var player Player

func SubEmpty[T comparable](value, defaultValue T) T {
	var zero T
	if value == zero {
		return defaultValue
	}
	return value
}

func initGame() {
	rooms = map[string]*Room{
		"кухня": NewRoom(
			"кухня",
			"ты находишься на кухне, ",
			"кухня, ничего интересного.",
			[]string{"коридор"},
			map[string][]string{"на столе": {"чай"}},
		),
		"комната": NewRoom(
			"комната",
			"",
			"ты в своей комнате.",
			[]string{"коридор"},
			map[string][]string{
				"на стуле": {"рюкзак"},
				"на столе": {"ключи", "конспекты"}},
		),
		"коридор": NewRoom(
			"коридор",
			"",
			"ничего интересного.",
			[]string{"кухня", "комната", "улица"},
			map[string][]string{},
		),
		"улица": NewRoom(
			"улица",
			"",
			"на улице весна.",
			[]string{"домой"},
			map[string][]string{},
		),
	}

	rooms["домой"] = rooms["коридор"]
	rooms["улица"].isOpen = false
	rooms["кухня"].taskDescription = "надо собрать рюкзак и идти в универ. "

	player = Player{
		room:  rooms["кухня"],
		items: []string{},
	}

	go func() {

	}()

}

func updateWorld() {
	var bag = false

	if player.room.name == "улица" {
		return
	}
	if !bag && player.isbag {
		bag = true
		rooms["кухня"].taskDescription = "надо идти в универ. "
	}
}

func handleCommand(command string) string {

	var result = ""

	args := strings.Split(command, " ")
	switch args[0] {
	case "осмотреться":
		result = player.lookAround()
	case "идти":
		result = player.goTo(args[1])
	case "надеть":
		switch args[1] {
		case "рюкзак":
			result = player.getBag()
		default:
			result = "нет такого"
		}
	case "взять":
		if len(args) < 2 {
			result = "не выбран предмет для взятия"
		}
		if player.isbag {
			result = player.getInBag(args[1])
		} else {
			result = "некуда класть"
		}
	case "применить":
		if len(args) < 3 {
			result = "не к чему применить"
		}
		result = player.activeItem(args[1], args[2])
	default:
		result = "неизвестная команда"
	}

	updateWorld()
	return result
}
