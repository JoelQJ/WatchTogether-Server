package server

import (
	"WatchTogether-Server/broadcast"
	"WatchTogether-Server/distpacher"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"strconv"
)



func Start(){
	var ipInterface string = "0.0.0.0"
	var port int = 1411
	var address string = ipInterface + ":" + strconv.Itoa(port)
	server, err := net.Listen("tcp", address)
	
	if err != nil {
		panic(err)
	}

	fmt.Println("Server abierto en el puerto " + address)
	defer func(server net.Listener) {
		err := server.Close()
		if err != nil {

		}
	}(server)

	for {
		conn, err := server.Accept()
		fmt.Println("Cliente Conectado")
		if err != nil {
			panic(err)
		}
		broadcast.AddConnection(conn)
		
		go handleConnection(conn)
	}
}



func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {

		sizeBuff := make([]byte, 4)

		_, err := io.ReadFull(conn, sizeBuff)
		if err != nil {
			broadcast.HandleDisconnect(conn)
			return
		}
		var packetSize int32
		binary.Read(bytes.NewBuffer(sizeBuff), binary.BigEndian, &packetSize)

		packetBuffer := make([]byte, packetSize)
		_, err = io.ReadFull(conn, packetBuffer)
		if err != nil {
			broadcast.HandleDisconnect(conn)
			return
		}
		fmt.Printf("Size:%d Payload:%v\n", packetSize, packetBuffer)
		distpacher.Dispatch(conn, bytes.NewBuffer(packetBuffer))
	}
}

