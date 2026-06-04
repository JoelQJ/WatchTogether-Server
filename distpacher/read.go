package distpacher

import (
	"WatchTogether-Server/broadcast"
	"WatchTogether-Server/events"
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
	register(HandShake, handleHandShake)
}

func Dispatch(client net.Conn, buff *bytes.Buffer) {
	var id PacketsIDS	
	err := binary.Read(buff, byteOrder, &id)
	if err != nil {
		fmt.Println("Error leyendo id del packet:", err)
		return
	}
	var isValidate bool = broadcast.IsValidated(client)
	if !isValidate && id != HandShake {
		broadcast.HandleDisconnect(client)
		return
	}
	
	funcion, ok := packets[int32(id)]
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
	events.FireVideoFinish(client)
}

func handleVideoLoaded(client net.Conn, buff *bytes.Buffer) {
	fmt.Println("Video Cargado")
}

func handleHandShake(client net.Conn, buff *bytes.Buffer){
	fmt.Println("Cliente Conectado")
	broadcast.ValidateConnection(client)

}
