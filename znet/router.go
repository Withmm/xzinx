package znet

import (
	"fmt"
	"zinx/ziface"
)

type BaseRouter struct{}

func (r *BaseRouter) PreHandle(req ziface.IRequest) {
	fmt.Println("pre...")
}
func (r *BaseRouter) Handle(req ziface.IRequest) {
	fmt.Println("ing...")
}
func (r *BaseRouter) PostHandle(req ziface.IRequest) {
	fmt.Println("post...")
}

func NewRouter() ziface.IRouter {
	return &BaseRouter{}
}
