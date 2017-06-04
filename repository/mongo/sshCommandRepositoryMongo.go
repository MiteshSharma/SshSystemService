package mongo

import (
	"github.com/MiteshSharma/SshSystemSetup/modal"
	"github.com/MiteshSharma/SshSystemSetup/repository/mongo/database"
	"gopkg.in/mgo.v2/bson"
)

type SshCommandRepositoryMongo struct {
}

func NewSshCommandRepositoryMongo() *SshCommandRepositoryMongo {
	sshCommandRepo:= &SshCommandRepositoryMongo{}
	return sshCommandRepo;
}

func CreateSshCommand(sshCommand *modal.SSHCommand) error {
	db := database.GetDatabaseManager()
	if err := db.Create("ssh_commands", sshCommand); err != nil {
		return err
	}

	return nil
}

func GetSshCommands() ([]modal.SSHCommand, error) {
	db := database.GetDatabaseManager()

	result := []modal.SSHCommand{}
	query := &bson.M{}
	if err := db.GetAllByQuery("ssh_commands", query, &result); err != nil {
		return result, err
	}

	return result, nil
}

func GetSshCommand(id string) ([]modal.SSHCommand, error) {
	db := database.GetDatabaseManager()

	sshCommand := []modal.SSHCommand{}
	if err := db.Get("ssh_commands", id, &sshCommand); err != nil {
		return sshCommand, err
	}

	return sshCommand, nil
}

func UpdateSshCommand(sshCommand *modal.SSHCommand) error {
	db := database.GetDatabaseManager()
	if err := db.Save("ssh_commands", sshCommand.Id, sshCommand); err != nil {
		return err
	}

	return nil
}

func CreateSshCommandStatus(sshCommandStatus *modal.SSHCommandStatus) error {
	db := database.GetDatabaseManager()
	if err := db.Create("ssh_command_statuses", sshCommandStatus); err != nil {
		return err
	}

	return nil
}

func GetSshCommandStatus(id string) (*modal.SSHCommandStatus, error) {
	db := database.GetDatabaseManager()

	sshCommand := []modal.SSHCommandStatus{}
	if err := db.Get("ssh_command_statuses", id, &sshCommand); err != nil {
		return sshCommand, err
	}

	return sshCommand, nil
}

func UpdateSshCommandStatus(sshCommandStatus *modal.SSHCommandStatus) error {
	db := database.GetDatabaseManager()
	if err := db.Save("ssh_command_statuses", sshCommandStatus.Id, sshCommandStatus); err != nil {
		return err
	}

	return nil
}