package repository

import (
	"github.com/MiteshSharma/SshSystemSetup/modal"
	"github.com/MiteshSharma/SshSystemSetup/repository/aws"
)

type InstanceRepository interface {
	GetDetails() ([]modal.InstanceDetail, error)
}

func GetInstanceRepository(platform string) InstanceRepository {
	return repository.NewEC2InstanceRepository();
}