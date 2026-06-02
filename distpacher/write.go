package distpacher

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func Write(buffer *bytes.Buffer, data any) {
	if err := binary.Write(buffer, byteOrder, data); err != nil {
		fmt.Println(err)
	}
}

func WriteString(buffer *bytes.Buffer, text string) {
	var size int32 = int32(len(text))
	Write(buffer, size)
	var textBytes []byte = []byte(text)
	Write(buffer, textBytes)

}

func CreateBufferForWritePacket(id PacketsIDS) *bytes.Buffer {
	var buffer *bytes.Buffer = new(bytes.Buffer)
	Write(buffer, id)
	return buffer
}


//Packets!
func WriteSetVideoPacket(url string) *bytes.Buffer {
	var buffer = CreateBufferForWritePacket(SetVideo)
	WriteString(buffer, url)
	return buffer
}

func WritePlayPusePacket(play bool) *bytes.Buffer{
	var buffer = CreateBufferForWritePacket(PlayPause)
	Write(buffer, play)
	return buffer
}

func WriteSetTimePacket(time float64)  *bytes.Buffer{
	var buffer = CreateBufferForWritePacket(SetTime)
	Write(buffer, time)

	return buffer
}