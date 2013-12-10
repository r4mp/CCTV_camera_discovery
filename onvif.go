package main

import (
	"github.com/r4mp/CCTV_camera_discovery/configuration"
	"log"
)

func main() {

	err := configuration.Load()
	configuration.Config.DiscoveryAddress = "239.255.255.250:3702"
	configuration.Save()

	if err != nil {
		log.Fatal(err)
		return
	}
}
