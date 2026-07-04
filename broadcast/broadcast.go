package broadcast

import (
	"WatchTogether-Server/events"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"slices"
	"sync"
)

var connections []net.Conn = make([]net.Conn, 0)
var validConnections []net.Conn = make([]net.Conn, 0)
var mu sync.Mutex

func AddConnection(conn net.Conn) {
	mu.Lock()
	connections = append(connections, conn)
	mu.Unlock()
}

func ValidateConnection(conn net.Conn) {
	mu.Lock()
	validConnections = append(validConnections, conn)
	mu.Unlock()
}

func IsValidated(conn net.Conn) bool {
	mu.Lock()
	defer mu.Unlock()
	return slices.Contains(validConnections, conn)
}

func HandleDisconnect(conn net.Conn) {
	fmt.Println("Cliente Desconectado")
	mu.Lock()
	connections = slices.DeleteFunc(connections, func(connIt net.Conn) bool {
		return connIt == conn
	})
	validConnections = slices.DeleteFunc(validConnections, func(connIt net.Conn) bool {
		return connIt == conn
	})
	mu.Unlock()
	events.FireClientDisconnect(conn)
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
	snapshot := make([]net.Conn, len(validConnections))
	copy(snapshot, validConnections)
	mu.Unlock()
	for _, client := range snapshot {
		if client == clientExcluded {
			continue
		}
		Send(client, buffer)
	}
}

func SenAll(buffer *bytes.Buffer) {
	SendBroadcast(nil, buffer)
}

func GetConnections() []net.Conn {
	mu.Lock()
	snapshot := make([]net.Conn, len(validConnections))
	copy(snapshot, validConnections)
	mu.Unlock()
	return snapshot
}
