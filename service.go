package task

import "github.com/goantor/deamon"

type RegisterHandler func(manager *deamon.DefaultDaemonManager)

type Task struct{}

func (t *Task) Register(handler RegisterHandler) {
	handler(deamon.DaemonManager)
}

func (t *Task) Boot() {
	deamon.DaemonManager.Start()
}
