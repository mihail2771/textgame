package main

import (
	"fmt"
	"strings"
)

type Room struct {
	ID              RoomID              `json:"id"`
	Name            string              `json:"name"`
	Description     string              `json:"description"`
	InDescription   string              `json:"inDescription"`
	Exits           []RoomID            `json:"exits"`
	Items           map[Position][]Item `json:"items"`
	TaskDescription string              `json:"taskDescription,omitempty"`
	IsOpen          bool                `json:"isOpen,omitempty"`
}

func (r *Room) lookAround() string {
	return fmt.Sprintf("%s%s%s"+msgCanGo, r.Description, r.mapToString(), r.TaskDescription, formatRoomList(r.Exits))
}

func (r *Room) lookAroundIn() string {
	return fmt.Sprintf(msgRoomIn, r.InDescription, formatRoomList(r.Exits))
}

func (r *Room) mapToString() string {
	var parts []string
	for _, position := range []Position{PositionTable, PositionChair} {
		itemIDs := r.Items[position]
		if len(itemIDs) == 0 {
			continue
		}
		parts = append(parts, fmt.Sprintf("%s: %s", positionLabels[position], joinItemLabels(itemIDs)))
	}
	if len(parts) == 0 {
		return msgRoomEmpty
	}
	if r.TaskDescription != "" {
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
		ID:              id,
		Name:            name,
		Description:     description,
		InDescription:   inDescription,
		Items:           items,
		Exits:           exits,
		IsOpen:          true,
		TaskDescription: "",
	}
}
