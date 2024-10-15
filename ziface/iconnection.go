package ziface

import "net"

type IConnection interface {
	Run()

	GetTCPConnection() *net.TCPConn

	RemoteAddr() net.Addr
}
