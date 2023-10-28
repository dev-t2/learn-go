package main

import "learn-go/12-interfaces/gadget"

func playList(device gadget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}

	device.Stop()
}

func main() {
	player := gadget.TapePlayer{}
	// player := gadget.TapeRecorder{}

	mixtape := []string{"Remember Me", "Feel My Rhythm", "Dumb Dumb", "Bad Boy"}

	playList(player, mixtape)
}