package worker

import (
	"time"
	"github.com/MiteshSharma/SshSystemSetup/taskExecutor/work"
	"github.com/MiteshSharma/SshSystemSetup/taskExecutor/taskHandler"
)

type Worker struct  {
	Id string
	Work chan work.Work
	WorkerQueue chan Worker
	workResponse chan taskHandler.WorkResponse
	Quit	chan bool
}

func NewWorker(id string, taskWorkerQueue chan Worker, communication chan taskHandler.WorkResponse) *Worker  {
	worker := &Worker{
		Id: id,
		Work: make(chan work.Work),
		WorkerQueue: taskWorkerQueue,
		workResponse: communication,
		Quit: make(chan bool)}
	return worker
}

func (w *Worker) Start()  {
	go func() {
		for {
			// Adding worker in worker queue
			w.WorkerQueue <- *w
			select {
			case task := <- w.Work:
				response := task.Execute()
				w.workResponse <- taskHandler.WorkResponse{Work: task, IsSuccess: response}
				time.Sleep(1 * time.Second)
			case <- w.Quit:
				// Stop this worker
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.Quit <- true
	}()
}