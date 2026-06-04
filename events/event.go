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

