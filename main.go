package main

import (
	"fmt"
	"zinx/znet"
)

func main() {
	fmt.Println("hello, zinx!")
	s := znet.NewServer("zinx[v0.1]")
	s.Run()
}
