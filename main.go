package main

import (
	"github.com/Noofbiz/wormWars1/astiGUI"
	"github.com/Noofbiz/wormWars1/astiServer"
)

func main() {
	astiServer.StartServer()

	astiGUI.StartGUI()
}
