package repository

import (
	"github.com/MiteshSharma/SshSystemSetup/modal"
	"github.com/MiteshSharma/SshSystemSetup/repository/mongo"
)

type SshCommandRepository interface {
	CreateSshCommand(*modal.SSHCommand) error
	GetSshCommands() ([]modal.SSHCommand, error)
	GetSshCommand(id string) ([]modal.SSHCommand, error)
	UpdateSshCommand(*modal.SSHCommand) error
	CreateSshCommandStatus(*modal.SSHCommandStatus) error
	GetSshCommandStatus(id string) (*modal.SSHCommandStatus, error)
	UpdateSshCommandStatus(*modal.SSHCommandStatus) error
}

func GetSshCommandRepository() *SshCommandRepository {
	return mongo.NewSshCommandRepositoryMongo();
}
