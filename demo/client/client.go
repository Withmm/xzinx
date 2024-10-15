package main

import (
	"fmt"
	"net"
	"time"
	"zinx/utils"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8089")
	utils.HandleError(err)
	for {
		conn.Write([]byte("hello, I'm client1."))
		buf := make([]byte, 256)
		conn.Read(buf)
		fmt.Printf("%s", buf)
		time.Sleep(1 * time.Second)
	}
}
