package task

import (
	"github.com/goantor/deamon"
)

const (
	Once = iota + 1
	Loop
	Time
)

type RegisterHandler func()

type ITask interface {
}

type Task struct {
	RouteHandler RegisterHandler
}

//func (t *Task) Register(handler RegisterHandler) {
//	handler(deamon.DaemonManager)
//}

func (t *Task) Boot() error {
	//c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Interrupt, os.Kill)

	go deamon.Start()

	//<-c
	return nil
}
