package main

import (
	"fmt"
	"strings"
)

type Room struct {
	id              RoomID
	name            string
	description     string
	taskDescription string
	inDescription   string
	items           map[Position][]Item
	exits           []RoomID
	isOpen          bool
}

func (r *Room) lookAround() string {
	return fmt.Sprintf("%s%s%s"+msgCanGo, r.description, r.mapToString(), r.taskDescription, formatRoomList(r.exits))
}

func (r *Room) lookAroundIn() string {
	return fmt.Sprintf(msgRoomIn, r.inDescription, formatRoomList(r.exits))
}

func (r *Room) mapToString() string {
	var parts []string
	for _, position := range positions {
		itemIDs := r.items[position]
		if len(itemIDs) == 0 {
			continue
		}
		parts = append(parts, fmt.Sprintf("%s: %s", positionLabels[position], joinItemLabels(itemIDs)))
	}
	if len(parts) == 0 {
		return msgRoomEmpty
	}
	if r.taskDescription != "" {
		return fmt.Sprintf("%s, ", strings.Join(parts, ", "))
	}
	return fmt.Sprintf("%s. ", strings.Join(parts, ", "))
}

func joinItemLabels(items []Item) string {
	var labels []string
	for _, item := range items {
		labels = append(labels, itemLabels[item])
	}
	return strings.Join(labels, ", ")
}

func formatRoomList(roomIDs []RoomID) string {
	var names []string
	for _, id := range roomIDs {
		names = append(names, roomLabels[id])
	}
	return strings.Join(names, ", ")
}

func NewRoom(id RoomID, name string, description string, inDescription string, exits []RoomID, items map[Position][]Item) *Room {
	return &Room{
		id:              id,
		name:            name,
		description:     description,
		inDescription:   inDescription,
		items:           items,
		exits:           exits,
		isOpen:          true,
		taskDescription: "",
	}
}
