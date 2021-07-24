package main

import "github.com/headfirstgo/gadget"

type Player interface {
	Play(string)
	Stop()
}

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	recoder, ok := player.(gadget.TapeRecorder)
	if ok {
		recoder.Record()
	}
}

func main() {
	TryOut(gadget.TapeRecorder{})
}
