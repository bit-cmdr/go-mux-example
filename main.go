package main

import "github.com/bit-cmdr/go-mux-example/app"

func main() {
	app := &app.App{}
	app.Initialize()
	app.Run(":3000")
}
