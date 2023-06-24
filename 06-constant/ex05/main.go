package main

import "fmt"

type Room uint8


// Bit Flags
const (
	MasterRoom Room = 1 << iota
	BathRoom
	Kitchen
	LivingRoom
)

func SetLight(rooms, room Room) Room {
	return rooms | room
}

func ResetLight(rooms, room Room) Room {
	return rooms &^ room
}

func IsTurnOn(rooms, room Room) bool {
	return (rooms & room) == room
}

func TurnOnLights(rooms Room) {
	if IsTurnOn(rooms, MasterRoom) {
		fmt.Println("TurnOn MasterRoom Light")
	}

	if IsTurnOn(rooms, BathRoom) {
		fmt.Println("TurnOn BathRoom Light")
	}

	if IsTurnOn(rooms, Kitchen) {
		fmt.Println("TurnOn Kitchen Light")
	}

	if IsTurnOn(rooms, LivingRoom) {
		fmt.Println("TurnOn LivingRoom Light")
	}
}

func main() {
	var rooms Room = 0

	rooms = SetLight(rooms, MasterRoom)
	rooms = SetLight(rooms, LivingRoom)

	TurnOnLights(rooms)

	rooms = ResetLight(rooms, LivingRoom)

	TurnOnLights(rooms)
}