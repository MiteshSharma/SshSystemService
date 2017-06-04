package taskExecutor

import (
	"github.com/MiteshSharma/SshSystemSetup/taskExecutor/taskHandler"
	"github.com/MiteshSharma/SshSystemSetup/taskExecutor/dispatcher"
	"github.com/MiteshSharma/SshSystemSetup/taskExecutor/work"
)

var TaskExecutor *Executor

type Executor struct {
	numWorker int
}

func NewExecutor(numWorker int) *Executor  {
	executor := &Executor{
		numWorker: numWorker,
	}
	return executor
}

func Init() *Executor {
	TaskExecutor = NewExecutor(1)
	return TaskExecutor
}

var TaskHandler *taskHandler.TaskHandler
var Dispatcher *dispatcher.Dispatcher

var communicationChan chan taskHandler.WorkResponse

func (e *Executor) Start() {
	communicationChan = make(chan taskHandler.WorkResponse)
	// Start task handler
	TaskHandler = taskHandler.NewTaskHandler(communicationChan)
	TaskHandler.Start()

	// Start dispatcher, dispatcher will start needed workers
	Dispatcher = dispatcher.NewDispatcher(e.numWorker, communicationChan)
	Dispatcher.Start()
}

func (e *Executor) ExecuteTask(work work.Work)  {
	Dispatcher.AssignTask(work)
}

func (e *Executor) Stop() {
	// Stop task reader
	TaskHandler.Stop()

	// Stop dispatcher
	Dispatcher.Stop()
}
