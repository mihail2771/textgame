package main

type Position string

const (
	PositionTable Position = "table"
	PositionChair Position = "chair"
)

var positionLabels = map[Position]string{
	PositionTable: "на столе",
	PositionChair: "на стуле",
}

type Item string

const (
	ItemTea      Item = "tea"
	ItemBackpack Item = "backpack"
	ItemKeys     Item = "keys"
	ItemNotes    Item = "notes"
)

var itemLabels = map[Item]string{
	ItemTea:      "чай",
	ItemBackpack: "рюкзак",
	ItemKeys:     "ключи",
	ItemNotes:    "конспекты",
}

var itemLabelToID = map[string]Item{
	"чай":       ItemTea,
	"рюкзак":    ItemBackpack,
	"ключи":     ItemKeys,
	"конспекты": ItemNotes,
}

type RoomID string

const (
	RoomKitchen RoomID = "kitchen"
	RoomRoom    RoomID = "room"
	RoomHallway RoomID = "hallway"
	RoomStreet  RoomID = "street"
	RoomHome    RoomID = "home"
)

var roomLabels = map[RoomID]string{
	RoomKitchen: "кухня",
	RoomRoom:    "комната",
	RoomHallway: "коридор",
	RoomStreet:  "улица",
	RoomHome:    "домой",
}

var roomLabelToID = map[string]RoomID{
	"кухня":   RoomKitchen,
	"комната": RoomRoom,
	"коридор": RoomHallway,
	"улица":   RoomStreet,
	"домой":   RoomHome,
}

const (
	cmdLook  = "осмотреться"
	cmdGo    = "идти"
	cmdWear  = "надеть"
	cmdTake  = "взять"
	cmdApply = "применить"

	msgUnknownCommand = "неизвестная команда"
	msgNoPath         = "не выбран путь"
	msgNoSuch         = "нет такого"
	msgNoWear         = "не выбрано что надеть"
	msgNoTake         = "не выбран предмет для взятия"
	msgNoBagSpace     = "некуда класть"
	msgNoApply        = "не к чему применить"
	msgAlreadyWearing = "вы уже надели рюкзак"
	msgWoreBackpack   = "вы надели: рюкзак"
	msgDoorClosed     = "дверь закрыта"
	msgDoorOpened     = "дверь открыта"
	msgRoomEmpty      = "пустая комната. "
	msgCanGo          = "можно пройти - %s"
	msgRoomIn         = "%s можно пройти - %s"
	subjectDoor       = "дверь"
	taskNeedBackpack  = "надо собрать рюкзак и идти в универ. "
	taskGoUniversity  = "надо идти в универ. "
	descKitchenStart  = "ты находишься на кухне, "
	descKitchenIn     = "кухня, ничего интересного."
	descYourRoomIn    = "ты в своей комнате."
	descHallwayIn     = "ничего интересного."
	descStreetIn      = "на улице весна."
)

func msgItemAdded(item Item) string {
	return "предмет добавлен в инвентарь: " + itemLabels[item]
}

func msgNoItemInInventory(label string) string {
	return "нет предмета в инвентаре - " + label
}

func msgNoPathToRoom(room RoomID) string {
	return "нет пути в " + roomLabels[room]
}
