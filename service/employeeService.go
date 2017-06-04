package service

import (
	"github.com/MiteshSharma/SshSystemSetup/repository"
	"github.com/MiteshSharma/SshSystemSetup/modal"
	"github.com/satori/go.uuid"
)

type EmployeeService struct {
	employeeRepo repository.EmployeeRepository
}

func NewEmployeeService() *EmployeeService {
	employeeService:= &EmployeeService{}
	employeeService.employeeRepo = repository.GetEmployeeRepository()
	return employeeService;
}

func (scs *EmployeeService) CreateEmployee(employee *modal.Employee) (*modal.Employee, error) {
	if employee.Id == "" {
		employee.Id = uuid.NewV4().String()
	}
	err := scs.employeeRepo.CreateEmployee(employee)
	if (err != nil) {
		return nil, err
	}
	return employee, nil
}

func (scs *EmployeeService) UpdateEmployee(employee *modal.Employee) (*modal.Employee, error) {
	err := scs.employeeRepo.UpdateEmployee(employee)
	if (err != nil) {
		return nil, err
	}
	return employee, nil
}

func (scs *EmployeeService) GetEmployees() ([]modal.Employee, error) {
	employees, err := scs.employeeRepo.GetEmployees();
	if (err != nil) {
		return nil, err
	}
	return employees, nil
}