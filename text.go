package main

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
