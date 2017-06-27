package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Noofbiz/wormWars1/astiGUI"
	"github.com/Noofbiz/wormWars1/astiServer"
)

func main() {
	astiServer.StartServer()

	astiGUI.StartGUI()

	cleanup()
}

func cleanup() {
	var d *os.File
	var err error
	if d, err = os.Open("astiServer/static/pieCharts"); err != nil {
		log.Fatalf("Unable to open pieCharts folder. Error: %v", err.Error())
	}
	defer d.Close()

	var names []string
	if names, err = d.Readdirnames(-1); err != nil {
		log.Fatalf("Unable to read pieCharts files. Error: %v", err.Error())
	}
	for _, name := range names {
		if err = os.RemoveAll(filepath.Join("astiServer/static/pieCharts", name)); err != nil {
			log.Fatalf("Unable to remove all files from pieCharts folder. Error: %v", err.Error())
		}
	}
}
