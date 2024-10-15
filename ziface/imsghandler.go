package ziface

type IMsgHandler interface {
	AddRouter(n int64, r IRouter)

	DoMsgHandle(req IRequest)
}
