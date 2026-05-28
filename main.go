package main

import (
	"WatchTogether-Server/packets"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"strconv"
)

var connections []net.Conn = make([]net.Conn, 0)

func main() {

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

		connections = append(connections, conn)
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {

		sizeBuff := make([]byte, 4)
		_, err := io.ReadFull(conn, sizeBuff)
		if err != nil {
			return
		}

		var packetSize int32 = int32(binary.BigEndian.Uint32(sizeBuff))

		packetBuffer := make([]byte, packetSize)
		_, err = io.ReadFull(conn, packetBuffer)
		if err != nil {
			fmt.Println("Cliente desconectado ")
		}
		packets.Dispatch(conn, bytes.NewBuffer(packetBuffer))
	}
}
