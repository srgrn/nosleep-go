package main

import (
	// "github.com/go-vgo/robotgo"
	"github.com/ColdSauce/Go-Cheese"
	"time"
)

func main() {
	tick := time.Tick(120 * time.Second)
	for {
		select {
		case <-tick:
			current, _ := Gotem.GetMousePosition()
			current.X = current.X + 1
			Gotem.MoveMouseTo(current)
			current.X = current.X - 1
			// time.Sleep(200 * time.Millisecond)
			Gotem.MoveMouseTo(current)
			// x, y := robotgo.GetMousePos()
			// robotgo.MoveMouseSmooth(x+1, y-1, 1.0, 100.0)
			// robotgo.MoveMouseSmooth(x-1, y+1, 1.0, 100.0)
		}
	}

}
