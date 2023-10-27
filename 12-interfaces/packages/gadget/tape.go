package gadget

import "fmt"

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopped")
}

type TapeRecorder struct {
	MicroPhones int
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapeRecorder) Record(song string) {
	fmt.Println("Recording", song)
}

func (t TapeRecorder) Stop() {
	fmt.Println("Stopped")
}