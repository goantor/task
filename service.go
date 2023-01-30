package task

import "github.com/goantor/deamon"

const (
	Once = iota + 1
	Loop
	Time
)

type RegisterHandler func(manager *deamon.DefaultDaemonManager)

type ITask interface {
	Handler() deamon.TaskFunc
}

type Task struct{}

func (t *Task) Register(handler RegisterHandler) {
	handler(deamon.DaemonManager)
}

func (t *Task) Boot() {
	deamon.DaemonManager.Start()
}
