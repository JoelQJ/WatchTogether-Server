package distpacher

import (
	"WatchTogether-Server/broadcast"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

type Packet func(client net.Conn, data *bytes.Buffer)

var packets map[int32]Packet = make(map[int32]Packet)

func register(id PacketsIDS, packet Packet) {
	packets[int32(id)] = packet
}

func RegisterPackets() {
	register(Ping, handlePing)
	register(PlayPause, handlePlayPause)
	register(SetTime, handleSetTime)
	register(VideoFinish, handleVideoFinish)
	register(VideoLoaded, handleVideoLoaded)
}

func Dispatch(client net.Conn, buff *bytes.Buffer) {
	var id int32
	err := binary.Read(buff, byteOrder, &id)
	if err != nil {
		fmt.Println("Error leyendo id del packet:", err)
		return
	}
	funcion, ok := packets[id]
	if ok {
		funcion(client, buff)
	}
}

// Packets!
func handlePing(client net.Conn, buff *bytes.Buffer) {
	var buffer = CreateBufferForWritePacket(Ping)
	buffer.Write(buff.Bytes())
	broadcast.Send(client, buffer)
	fmt.Println("Ping! Enviando Pong!")
}

func handlePlayPause(client net.Conn, buff *bytes.Buffer) {
	var buffer = CreateBufferForWritePacket(PlayPause)
	buffer.Write(buff.Bytes())
	broadcast.SenAll(buffer)
}

func handleSetTime(client net.Conn, buff *bytes.Buffer) {
	var buffer = CreateBufferForWritePacket(SetTime)
	buffer.Write(buff.Bytes())
	broadcast.SenAll(buffer)
}

func handleVideoFinish(client net.Conn, buff *bytes.Buffer) {
	fmt.Println("Video terminado")
}

func handleVideoLoaded(client net.Conn, buff *bytes.Buffer) {
	fmt.Println("Video Cargado")
}
