package distpacher

import (
	"WatchTogether-Server/broadcast"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type Packet func(client net.Conn, data *bytes.Buffer)
var packets map[int]Packet = make(map[int]Packet)

func register(id int, packet Packet) {
	packets[id] = packet
}

func RegisterPackets(){
	register(0, handlePing)
	register(1, handlePlayPause)
	register(3, handleSetTime)

 go func() {
        time.Sleep(10 * time.Second)
        broadcast.SendBroadcast(nil, WriteSetVideoPacket("otaku.mkv"))

        go func(){
        	time.Sleep(10 * time.Second)
        broadcast.SendBroadcast(nil, WriteSetVideoPacket("otaku.mkv"))
        }()
    }()
}

func Dispatch(client net.Conn, buff *bytes.Buffer) {
	var id int32
	err := binary.Read(buff, byteOrder, &id)
	if err != nil {
		fmt.Println("Error leyendo id del packet:", err)
		return
	}
	funcion, ok := packets[int(id)]
	if ok {
		funcion(client, buff)
	}
}

func handlePing(client net.Conn, buff *bytes.Buffer){
	buffer := CreateBufferForWritePacket(0)
	buffer.Write(buff.Bytes())
	broadcast.Send(client, buffer)
	fmt.Println("Ping! Enviando Pong!")
}

func handlePlayPause(client net.Conn, buff *bytes.Buffer){
	buffer := CreateBufferForWritePacket(1)
	buffer.Write(buff.Bytes())
	broadcast.SendBroadcast(nil, buffer)
}

func handleSetTime(client net.Conn, buff *bytes.Buffer){
	buffer := CreateBufferForWritePacket(3)
	buffer.Write(buff.Bytes())
	broadcast.SendBroadcast(nil, buffer)
}

