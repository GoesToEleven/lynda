package main

import (
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/touch"
	"log"
)

func main() {
	app.Main(func(a app.App) {
		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
			case touch.Event:
				touchX := e.X
				touchY := e.Y
				log.Println("Mouse event:", touchX, touchY)
			}
		}
	})
}