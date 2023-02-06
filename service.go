package task

import (
	"github.com/goantor/deamon"
	"os"
	"os/signal"
)

const (
	Once = iota + 1
	Loop
	Time
)

type RegisterHandler func(manager *deamon.DefaultDaemonManager)

type ITask interface {
	Handler() deamon.TaskFunc
}

type Task struct {
	RouteHandler RegisterHandler
}

//func (t *Task) Register(handler RegisterHandler) {
//	handler(deamon.DaemonManager)
//}

func (t *Task) Boot() error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	t.RouteHandler(deamon.DaemonManager)
	deamon.DaemonManager.Start()

	<-c
	return nil
}
