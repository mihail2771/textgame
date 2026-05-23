package main

type Player struct {
	room  *Room
	items []string
	isbag bool
}

func (p *Player) getBag() string {
	if p.isbag {
		return "вы уже надели рюкзак"
	}
	for _, position := range positions {
		for i, item := range p.room.items[position] {
			if item == "рюкзак" {
				p.room.items[position] = append(p.room.items[position][:i], p.room.items[position][i+1:]...)
				break
			}
		}
	}
	p.isbag = true
	return "вы надели: рюкзак"
}

func (p *Player) lookAround() string {
	return p.room.lookAround()
}

func (p *Player) getInBag(itemName string) string {
	for pos, itemsInPosition := range p.room.items {
		for i, item := range itemsInPosition {
			if item == itemName {
				p.items = append(p.items, itemName)
				p.room.items[pos] = append(p.room.items[pos][:i], p.room.items[pos][i+1:]...)
				if len(p.room.items[pos]) == 0 {
					delete(p.room.items, pos)
				}
				return "предмет добавлен в инвентарь: " + itemName
			}
		}
	}
	return "нет такого"
}

func (p *Player) goTo(roomName string) string {
	found := false
	for _, exit := range p.room.exits {
		if exit == roomName {
			found = true
			break
		}
	}

	if !found {
		return "нет пути в " + roomName
	}

	if !rooms[roomName].isOpen {
		return "дверь закрыта"
	}

	p.room = rooms[roomName]
	return p.room.lookAroundIn()
}

func (p *Player) activeItem(itemName string, subject string) string {
	itemInInventory := false
	for _, item := range p.items {
		if item == itemName {
			itemInInventory = true
			break
		}
	}
	if p.isbag && itemInInventory {
		if itemName == "ключи" && subject == "дверь" {
			rooms["улица"].isOpen = true
			return "дверь открыта"
		}
		return "не к чему применить"
	}
	return "нет предмета в инвентаре - " + itemName
}
