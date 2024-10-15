package ziface

type IServer interface {
	Run()

	RegisterRouter(n int64, r IRouter)
}
