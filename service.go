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
	Handler(handler deamon.Handler)
}

type Task struct{}

func (t *Task) Register(handler RegisterHandler) {
	handler(deamon.DaemonManager)
}

func (t *Task) Boot() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	deamon.DaemonManager.Start()

	<-c
}
