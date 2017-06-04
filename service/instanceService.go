package service

import (
	"github.com/MiteshSharma/SshSystemSetup/repository"
	"github.com/MiteshSharma/SshSystemSetup/modal"
)

type InstanceService struct {
	instRepo repository.InstanceRepository
}

func NewInstanceService(instanceRepository repository.InstanceRepository) *InstanceService {
	instanceService:= &InstanceService{}
	instanceService.instRepo = instanceRepository
	return instanceService;
}

func (instService *InstanceService) GetDetails() ([]modal.InstanceDetail, error) {
	return instService.instRepo.GetDetails();
}
