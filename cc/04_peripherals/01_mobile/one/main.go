package main

import (
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/exp/app/debug"
	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/gl"
	"golang.org/x/mobile/event/size"
)

func main() {
	app.Main(func(a app.App) {
		sz := size.Event{}
		ctx, _ := gl.NewContext()
		fps := debug.NewFPS(glutil.NewImages(ctx))
		fps.Draw(sz)
	})
}