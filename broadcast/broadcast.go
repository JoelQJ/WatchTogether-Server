package broadcast

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"slices"
	"sync"
)

var connections []net.Conn = make([]net.Conn, 0)
var mu sync.Mutex

func AddConnection(conn net.Conn) {
	mu.Lock()
	connections = append(connections, conn)
	mu.Unlock()
}

func HandleDisconnect(conn net.Conn) {
	fmt.Println("Cliente Desconectado")
	mu.Lock()
	connections = slices.DeleteFunc(connections, func(connIt net.Conn) bool {
		return connIt == conn
	})
	mu.Unlock()
	conn.Close()
}

func Send(client net.Conn, buffer *bytes.Buffer) {
	sizeBuff := new(bytes.Buffer)
	binary.Write(sizeBuff, binary.BigEndian, int32(buffer.Len()))

	buffers := net.Buffers{sizeBuff.Bytes(), buffer.Bytes()}
	buffers.WriteTo(client)
}

func SendBroadcast(clientExcluded net.Conn, buffer *bytes.Buffer) {
	mu.Lock()
	snapshot := make([]net.Conn, len(connections))
	copy(snapshot, connections)
	mu.Unlock()

	for _, client := range snapshot {
		if client == clientExcluded {
			continue
		}
		Send(client, buffer)
	}
}

func SenAll(buffer *bytes.Buffer){
	SendBroadcast(nil, buffer)
}
