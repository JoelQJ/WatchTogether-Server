package packets

import (
	"bytes"
	"encoding/binary"
	"net"
)

type Packet func(client net.Conn, data *bytes.Buffer)

var packets map[int]Packet = make(map[int]Packet)

func register(id int, packet Packet) {
	packets[id] = packet
}

func Dispatch(client net.Conn, buff *bytes.Buffer) {
	var id int32
	err := binary.Read(buff, binary.BigEndian, &id)
	if err != nil {
		panic(err)
	}
	funcion, ok := packets[int(id)]
	if ok {
		funcion(client, buff)
	}
}
