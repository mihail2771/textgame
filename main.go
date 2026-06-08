package main

import (
	"strings"
)

var positions = []Position{PositionTable, PositionChair}

var rooms map[RoomID]*Room
var player Player

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

func initGame() {
	rooms = map[RoomID]*Room{
		RoomKitchen: NewRoom(
			RoomKitchen,
			roomLabels[RoomKitchen],
			descKitchenStart,
			descKitchenIn,
			[]RoomID{RoomHallway},
			map[Position][]Item{PositionTable: {ItemTea}},
		),
		RoomRoom: NewRoom(
			RoomRoom,
			roomLabels[RoomRoom],
			"",
			descYourRoomIn,
			[]RoomID{RoomHallway},
			map[Position][]Item{
				PositionChair: {ItemBackpack},
				PositionTable: {ItemKeys, ItemNotes}},
		),
		RoomHallway: NewRoom(
			RoomHallway,
			roomLabels[RoomHallway],
			"",
			descHallwayIn,
			[]RoomID{RoomKitchen, RoomRoom, RoomStreet},
			map[Position][]Item{},
		),
		RoomStreet: NewRoom(
			RoomStreet,
			roomLabels[RoomStreet],
			"",
			descStreetIn,
			[]RoomID{RoomHome},
			map[Position][]Item{},
		),
	}

	rooms[RoomHome] = rooms[RoomHallway]
	rooms[RoomStreet].isOpen = false
	rooms[RoomKitchen].taskDescription = taskNeedBackpack

	player = Player{
		room:  rooms[RoomKitchen],
		items: []Item{},
	}
}

func updateWorld() {
	if player.room.id == RoomStreet {
		return
	}
	if player.isbag {
		rooms[RoomKitchen].taskDescription = taskGoUniversity
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
		result = player.lookAround()
	case cmdGo:
		if len(args) < 2 {
			return msgNoPath
		}
		roomID, ok := parseRoom(args[1])
		if !ok {
			return msgNoSuch
		}
		result = player.goTo(roomID)
	case cmdWear:
		if len(args) < 2 {
			return msgNoWear
		}
		itemID, ok := parseItem(args[1])
		if !ok || itemID != ItemBackpack {
			return msgNoSuch
		}
		result = player.getBag()
	case cmdTake:
		if len(args) < 2 {
			return msgNoTake
		}
		itemID, ok := parseItem(args[1])
		if !ok {
			return msgNoSuch
		}
		if player.isbag {
			result = player.getInBag(itemID)
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
			result = player.activeItem(itemID, args[2])
		}
	default:
		result = msgUnknownCommand
	}

	updateWorld()
	return result
}
