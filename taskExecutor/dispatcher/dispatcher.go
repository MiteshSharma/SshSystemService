package dispatcher

import (
	"github.com/satori/go.uuid"
	"github.com/MiteshSharma/SshSystemSetup/taskExecutor/taskHandler"
	"github.com/MiteshSharma/SshSystemSetup/taskExecutor/work"
	"github.com/MiteshSharma/SshSystemSetup/taskExecutor/worker"
)

type Dispatcher struct  {
	NumWorker int
	workResponse chan taskHandler.WorkResponse
	Quit	chan bool
}

func NewDispatcher(numWorker int, communication chan taskHandler.WorkResponse) *Dispatcher  {
	dispatcher := &Dispatcher{
		NumWorker: numWorker,
		workResponse: communication,
		Quit: make(chan bool)}
	return dispatcher
}

var TaskWorkerQueue chan worker.Worker

func (d *Dispatcher) Start()  {
	TaskWorkerQueue = make(chan worker.Worker, d.NumWorker)

	for count:= 0; count < d.NumWorker; count++ {
		worker := worker.NewWorker(uuid.NewV4().String(), TaskWorkerQueue, d.workResponse)
		worker.Start()
	}

	go func() {
		var work work.Work
		for {
			select {
			case work = <- taskHandler.TaskQueue:
				go func() {
					var worker worker.Worker = <-TaskWorkerQueue
					worker.Work <- work
				}()
			case <- d.Quit:
				return
			}
		}
	}()
}

func (d *Dispatcher) Stop() {
	go func() {
		d.Quit <- true
	}()
}

func (d *Dispatcher) AssignTask(work work.Work) {
	taskHandler.AssignWork(work)
}