package znet

import "zinx/ziface"

type MsgHandler struct {
	Apis map[int64]ziface.IRouter
}

func (mh *MsgHandler) AddRouter(n int64, r ziface.IRouter) {
	mh.Apis[n] = r
}

func (mh *MsgHandler) DoMsgHandle(req ziface.IRequest) {
	msgID := req.GetMsg().GetMsgId()

	router, ok := mh.Apis[int64(msgID)]
	if !ok {
		panic("No ")
		return
	}
	router.PreHandle(req)
	router.Handle(req)
	router.PostHandle(req)
}
func NewMsgHandler() ziface.IMsgHandler {
	mh := &MsgHandler{
		Apis: make(map[int64]ziface.IRouter),
	}
	return mh
}
