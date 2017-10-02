package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.LockOSThread()

	app := NewApplication("Minecraft")
	defer app.dispose()

	app.runLoop()
}

func debug(vs ...interface{}) {
	fmt.Println(vs...)
}
