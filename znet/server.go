package znet

import (
	"fmt"
	"net"
	"zinx/utils"
	"zinx/ziface"
)

const (
	BUFMAX = 256
)

type Server struct {
	Name string //server name

	IP string //ip

	Port string //port
}

func (s *Server) GetServerIp() string {
	return s.IP
}

func (s *Server) GetServerPort() string {
	return s.Port
}

func (s *Server) Run() {
	listener, err := net.Listen("tcp", s.IP+":"+s.Port)
	utils.HandleError(err)
	fmt.Printf("Server is listening on %s\n", listener.Addr().String())

	for {
		conn, err := listener.Accept()
		utils.HandleError(err)

		go func(conn net.Conn) {
			defer conn.Close()
			for {
				buf := make([]byte, BUFMAX)
				n, err := conn.Read(buf)
				utils.HandleError(err)
				conn.Write(buf[:n])
			}
		}(conn)
	}
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name: name,
		IP:   "127.0.0.1",
		Port: "8089",
	}
	return s
}
