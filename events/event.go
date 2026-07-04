package events

import "net"

var onVideoFinish []func(client net.Conn)
func SubscribeVideoFinish(f func(net.Conn)) {
	onVideoFinish = append(onVideoFinish, f)
}
func FireVideoFinish(client net.Conn) {
	for _, f := range onVideoFinish {
		f(client)
	}
}

var onClientDisconnect []func(client net.Conn)
func SubscribeClientDisconnect(f func(net.Conn)) {
	onClientDisconnect = append(onClientDisconnect, f)
}
func FireClientDisconnect(client net.Conn) {
	for _, f := range onClientDisconnect {
		f(client)
	}
}

