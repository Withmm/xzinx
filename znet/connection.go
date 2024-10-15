package znet

import (
	"fmt"
	"net"
	"zinx/utils"
	"zinx/ziface"
)

type Connection struct {
	Conn *net.TCPConn

	ConnID int // 连接的序号

	MsgHandler ziface.IMsgHandler
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Run() {
	// 获取处理信息的函数
	fmt.Printf("[ConnInfo]: conn:%d is running\n", c.ConnID)
	for {
		dp := NewDataPack()
		buf := make([]byte, 512)
		c.Conn.Read(buf)

		msg, err := dp.Unpack(buf)
		utils.HandleError(err)

		req := &Request{
			Conn: c,
			Msg:  msg,
		}
		c.MsgHandler.DoMsgHandle(req)
	}
}
