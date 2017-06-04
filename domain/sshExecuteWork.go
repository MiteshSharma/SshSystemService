package domain

import (
	"github.com/MiteshSharma/SshSystemSetup/modal"
	"bytes"
	"fmt"
	"github.com/MiteshSharma/SshSystemSetup/service"
)

type SshExecuteWork struct {
	taskId string
	client SshClient
	instance modal.InstanceDetail
	command modal.SSHCommand
}

func NewSshExecuteWork(client SshClient, inst modal.InstanceDetail,cmd modal.SSHCommand) *SshExecuteWork {
	work := &SshExecuteWork{
		client: client,
		instance: inst,
		command: cmd,
	}
	work.init()
	return work
}

func (sew *SshExecuteWork) init()  {
	// Create work in DB for tracking
	ssgCommandService := service.NewSshCommandService()
	commandStatus, err := ssgCommandService.CreateSshCommandStatus(sew.command, nil)
	if err != nil {
		return
	}
	sew.taskId = commandStatus.Id
}

func (sew *SshExecuteWork) GetId() string {
	return sew.taskId;
}

func (sew *SshExecuteWork) Execute() bool {
	var message bytes.Buffer
	sew.command.Stdout = &message

	err := sew.client.RunCommand(sew.instance, &sew.command);

	fmt.Print("Response is : "+ message.String())
	if err != nil {
		sew.onWorkComplete("Failed", message)
		return false
	}
	sew.onWorkComplete("Success", message)
	return true
}

func (sew *SshExecuteWork) onWorkComplete(status string, message string) {
	ssgCommandService := service.NewSshCommandService()
	ssgCommandService.UpdateSshCommandStatus(sew.taskId, status, message)
}
