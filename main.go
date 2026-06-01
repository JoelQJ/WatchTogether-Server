package main

import (
	"WatchTogether-Server/distpacher"
	"WatchTogether-Server/server"
)

func main() {
	distpacher.RegisterPackets()
	server.Start()

}
