package service

import (
	"github.com/satori/go.uuid"
	"github.com/MiteshSharma/SshSystemSetup/domain"
	"github.com/MiteshSharma/SshSystemSetup/modal"
	"github.com/MiteshSharma/SshSystemSetup/taskExecutor"
	"github.com/MiteshSharma/SshSystemSetup/repository"
)

type SshCommandObj struct {
	Command modal.SSHCommand
	Instance modal.InstanceDetail
}

type SshCommandService struct {
	sshCommandRepo repository.SshCommandRepository
}

func NewSshCommandService() *SshCommandService {
	sshCommandService:= &SshCommandService{}
	sshCommandService.sshCommandRepo = repository.GetSshCommandRepository()
	return sshCommandService;
}

func (scs *SshCommandService) RunCommand(cmdObj SshCommandObj) error {
	// Async task execution
	taskExecutor.TaskExecutor.ExecuteTask(domain.NewSshExecuteWork(domain.GetSshClient(), cmdObj.Instance, cmdObj.Command))
	return nil
}

func (scs *SshCommandService) CreateSshCommand(command *modal.SSHCommand) (*modal.SSHCommand, error) {
	if command.Id == "" {
		command.Id = uuid.NewV4().String()
	}
	err := scs.sshCommandRepo.CreateSshCommand(command)
	if (err != nil) {
		return nil, err
	}
	return command, nil
}

func (scs *SshCommandService) UpdateSshCommand(command *modal.SSHCommand) (*modal.SSHCommand, error) {
	err := scs.sshCommandRepo.UpdateSshCommand(command)
	if (err != nil) {
		return nil, err
	}
	return command, nil
}

func (scs *SshCommandService) GetSshCommands() ([]modal.SSHCommand, error) {
	commands, err := scs.sshCommandRepo.GetSshCommands()
	if (err != nil) {
		return nil, err
	}
	return commands, nil
}

func (scs *SshCommandService) CreateSshCommandStatus(command *modal.SSHCommand, user *modal.User) (*modal.SSHCommandStatus, error) {
	commandStatus := &modal.SSHCommandStatus{}
	commandStatus.Id = uuid.NewV4().String()
	commandStatus.CommandId = command.Id
	commandStatus.UserId = user.Id
	commandStatus.Status = "Start"

	err := scs.sshCommandRepo.CreateSshCommandStatus(commandStatus)
	if (err != nil) {
		return nil, err
	}
	return commandStatus, nil
}

func (scs *SshCommandService) UpdateSshCommandStatus(commandStatusId string, status string, message string) (*modal.SSHCommandStatus, error) {
	commandStatus, err := scs.sshCommandRepo.GetSshCommandStatus(commandStatusId)
	if (err != nil) {
		return nil, err
	}
	commandStatus.Status = status
	commandStatus.Message = message
	err = scs.sshCommandRepo.UpdateSshCommandStatus(commandStatus)
	if (err != nil) {
		return nil, err
	}
	return commandStatus, nil
}