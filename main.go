package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

type Router struct {
	znet.BaseRouter
}

var r Router

func (r *Router) PreHandle(req ziface.IRequest) {
	req.GetConn()
}
func (r *Router) Handle(req ziface.IRequest) {
	fmt.Println("ing...")
}
func (r *Router) PostHandle(req ziface.IRequest) {
	fmt.Println("post...")
}

func main() {
	fmt.Println("hello, zinx!")
	s := znet.NewServer("zinx[v0.1]")

	s.Run()
}
