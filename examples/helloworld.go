package main

import (
	"os"

	"github.com/special/qgoscene"
)

func main() {
	scene := qgoscene.NewScene("helloworld.qml", os.Args)
	scene.SetContextProperty("location", "World")
	os.Exit(scene.Exec())
}
