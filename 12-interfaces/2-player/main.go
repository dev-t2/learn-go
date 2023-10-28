package main

import (
	"fmt"
	"learn-go/12-interfaces/gadget"
)

type Player interface {
	Play(string)
	Stop()
}

func playList1(device gadget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}

	device.Stop()
}

func playList2(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}

	device.Stop()
}

func main() {
	player1 := gadget.TapePlayer{}
	// player1 := gadget.TapeRecorder{}

	mixtape := []string{"Remember Me", "Feel My Rhythm", "Dumb Dumb", "Bad Boy"}

	playList1(player1, mixtape)

	fmt.Println()

	var player2 Player = gadget.TapePlayer{}

	playList2(player2, mixtape)

	fmt.Println()

	player2 = gadget.TapeRecorder{}

	playList2(player2, mixtape)
}