package main

type Player struct {
	room  *Room
	items []Item
	isbag bool
}

func (p *Player) getBag() string {
	if p.isbag {
		return msgAlreadyWearing
	}
	for _, position := range positions {
		for i, item := range p.room.items[position] {
			if item == ItemBackpack {
				p.room.items[position] = append(p.room.items[position][:i], p.room.items[position][i+1:]...)
				break
			}
		}
	}
	p.isbag = true
	return msgWoreBackpack
}

func (p *Player) lookAround() string {
	return p.room.lookAround()
}

func (p *Player) getInBag(item Item) string {
	for pos, itemsInPosition := range p.room.items {
		for i, stored := range itemsInPosition {
			if stored == item {
				p.items = append(p.items, item)
				p.room.items[pos] = append(p.room.items[pos][:i], p.room.items[pos][i+1:]...)
				if len(p.room.items[pos]) == 0 {
					delete(p.room.items, pos)
				}
				return msgItemAdded(item)
			}
		}
	}
	return msgNoSuch
}

func (p *Player) goTo(roomID RoomID) string {
	found := false
	for _, exit := range p.room.exits {
		if exit == roomID {
			found = true
			break
		}
	}

	if !found {
		return msgNoPathToRoom(roomID)
	}

	if !rooms[roomID].isOpen {
		return msgDoorClosed
	}

	p.room = rooms[roomID]
	return p.room.lookAroundIn()
}

func (p *Player) activeItem(item Item, subject string) string {
	itemInInventory := false
	for _, stored := range p.items {
		if stored == item {
			itemInInventory = true
			break
		}
	}
	if p.isbag && itemInInventory {
		if item == ItemKeys && subject == subjectDoor {
			rooms[RoomStreet].isOpen = true
			return msgDoorOpened
		}
		return msgNoApply
	}
	return msgNoItemInInventory(itemLabels[item])
}
