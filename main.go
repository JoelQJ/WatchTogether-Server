package main

import (
	"WatchTogether-Server/console"
	"WatchTogether-Server/distpacher"
	"WatchTogether-Server/server"
)

func main() {
	distpacher.RegisterPackets()
	go server.Start()
	console.Start()


}
