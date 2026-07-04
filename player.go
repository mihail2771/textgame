package main

type Player struct {
	Room  *Room
	Items []Item
	Isbag bool
}

func (p *Player) getBag() string {
	if p.Isbag {
		return msgAlreadyWearing
	}
	for _, position := range []Position{PositionTable, PositionChair} {
		for i, item := range p.Room.Items[position] {
			if item == ItemBackpack {
				p.Room.Items[position] = append(p.Room.Items[position][:i], p.Room.Items[position][i+1:]...)
				break
			}
		}
	}
	p.Isbag = true
	return msgWoreBackpack
}

func (p *Player) lookAround() string {
	return p.Room.lookAround()
}

func (p *Player) getInBag(item Item) string {
	for pos, itemsInPosition := range p.Room.Items {
		for i, stored := range itemsInPosition {
			if stored == item {
				p.Items = append(p.Items, item)
				p.Room.Items[pos] = append(p.Room.Items[pos][:i], p.Room.Items[pos][i+1:]...)
				if len(p.Room.Items[pos]) == 0 {
					delete(p.Room.Items, pos)
				}
				return msgItemAdded(item)
			}
		}
	}
	return msgNoSuch
}

func (p *Player) goTo(roomID RoomID) string {
	found := false
	for _, exit := range p.Room.Exits {
		if exit == roomID {
			found = true
			break
		}
	}

	if !found {
		return msgNoPathToRoom(roomID)
	}

	if !game.Rooms[roomID].IsOpen {
		return msgDoorClosed
	}

	p.Room = game.Rooms[roomID]
	return p.Room.lookAroundIn()
}

func (p *Player) activeItem(item Item, subject string) string {
	itemInInventory := false
	for _, stored := range p.Items {
		if stored == item {
			itemInInventory = true
			break
		}
	}
	if p.Isbag && itemInInventory {
		if item == ItemKeys && subject == subjectDoor {
			game.Rooms[RoomStreet].IsOpen = true
			return msgDoorOpened
		}
		return msgNoApply
	}
	return msgNoItemInInventory(itemLabels[item])
}
