package repository

import (
	"github.com/MiteshSharma/SshSystemSetup/modal"
	"github.com/MiteshSharma/SshSystemSetup/repository/mongo"
)

type EmployeeRepository interface {
	GetEmployees() ([]modal.Employee, error)
	CreateEmployee(modal.Employee) error
	UpdateEmployee(modal.Employee) error
}

func GetEmployeeRepository() *EmployeeRepository {
	return mongo.NewEmployeeRepositoryMongo();
}