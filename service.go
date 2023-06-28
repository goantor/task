package task

import (
	"context"
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
	go deamon.Start()
	return nil
}

func (t *Task) Shutdown(ctx context.Context) error {
	return nil
}
