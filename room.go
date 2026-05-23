package main

import (
	"fmt"
	"strings"
)

type Room struct {
	name            string
	description     string
	taskDescription string
	inDescription   string
	items           map[string][]string
	exits           []string
	isOpen          bool
}

func (r *Room) lookAround() string {

	return fmt.Sprintf("%s%s%sможно пройти - %s", r.description, r.mapToString(), r.taskDescription, strings.Join(r.exits, ", "))
}

func (r *Room) lookAroundIn() string {
	return fmt.Sprintf("%s можно пройти - %s", r.inDescription, strings.Join(r.exits, ", "))
}

func (r *Room) mapToString() string {
	var parts []string
	for _, position := range positions {
		itemNames := r.items[position]
		if len(itemNames) == 0 {
			continue
		}
		parts = append(parts, fmt.Sprintf("%s: %s", position, strings.Join(itemNames, ", ")))

	}
	if len(parts) == 0 {
		return "пустая комната. "
	}
	if r.taskDescription != "" {
		return fmt.Sprintf("%s, ", strings.Join(parts, ", "))
	}
	return fmt.Sprintf("%s. ", strings.Join(parts, ", "))
}

func NewRoom(name string, description string, inDescription string, exits []string, items map[string][]string) *Room {
	return &Room{
		name:            name,
		description:     description,
		inDescription:   inDescription,
		items:           items,
		exits:           exits,
		isOpen:          true,
		taskDescription: "",
	}
}
