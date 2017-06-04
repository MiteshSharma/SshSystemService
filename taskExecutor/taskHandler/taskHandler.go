package taskHandler

import (
	"github.com/MiteshSharma/SshSystemSetup/taskExecutor/work"
	"fmt"
)

type TaskHandler struct  {
	response chan WorkResponse
	Quit chan bool
}

type WorkResponse struct  {
	Work work.Work
	IsSuccess bool
}

func NewTaskHandler(response chan WorkResponse) *TaskHandler {
	taskHandler := &TaskHandler{
		response: response,
		Quit: make(chan bool)}
	return taskHandler
}

var TaskQueue = make(chan work.Work, 100)

func (r *TaskHandler) Start()  {
	go func() {
		for {
			select {
			case taskResponse:= <-r.response:
				//logs.Logger.Debug("Response of task received with id : "+taskResponse.Work.GetId())
				fmt.Print("Task assigned"+ taskResponse.Work.GetId())
			}
		}
	}()
}

func AssignWork(work work.Work)  {
	TaskQueue <- work
}

func (r *TaskHandler) Stop()  {
	go func() {
		r.Quit <- true
	}()
}