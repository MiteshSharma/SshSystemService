package main

import (
	"runtime"
	"os"
	"os/signal"
	"syscall"
	"flag"
	"github.com/MiteshSharma/SshSystemSetup/utils"
	"github.com/MiteshSharma/SshSystemSetup/api"
	"github.com/MiteshSharma/SshSystemSetup/taskExecutor"
)

var configFileName string

func main() {
	parseCmdParams()

	utils.LoadConfig(configFileName)

	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)

	// Start task executor to run async tasks
	executor := taskExecutor.Init()
	executor.Start()
	// Init and start api server
	api.InitServer()
	api.StartServer()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-c

	api.StopServer()
	executor.Stop()
}

func parseCmdParams()  {
	flag.StringVar(&configFileName, "config", "config.json", "")
	flag.Parse()
}
