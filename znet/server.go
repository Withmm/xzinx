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

	IPVersion string // ip version

	IP   string //ip
	Port string //port

	MsgHandler ziface.IMsgHandler
}

func (s *Server) GetServerIp() string {
	return s.IP
}

func (s *Server) GetServerPort() string {
	return s.Port
}

func (s *Server) Run() {
	fmt.Println("[START] Server listening at IP:", s.IP, "Port:", s.Port)

	for {
		addr, err := net.ResolveTCPAddr(s.IPVersion, s.IP+":"+s.Port)
		utils.HandleError(err)

		listener, err := net.ListenTCP(s.IPVersion, addr)
		utils.HandleError(err)

		con, err := listener.AcceptTCP()
		Conn := &Connection{
			Conn:       con,
			ConnID:     0,
			MsgHandler: s.MsgHandler,
		}

		go Conn.Run()
	}
}

func (s *Server) RegisterRouter(n int64, r ziface.IRouter) {
	s.MsgHandler.AddRouter(n, r)
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:       name,
		IPVersion:  utils.GlobalObject.IPVersion,
		IP:         utils.GlobalObject.IP,
		Port:       utils.GlobalObject.Port,
		MsgHandler: NewMsgHandler(),
	}
	return s
}
