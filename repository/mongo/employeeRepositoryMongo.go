package mongo

import (
	"github.com/MiteshSharma/SshSystemSetup/repository/mongo/database"
	"github.com/MiteshSharma/SshSystemSetup/modal"
	"gopkg.in/mgo.v2/bson"
)

type EmployeeRepositoryMongo struct {
}

func NewEmployeeRepositoryMongo() *EmployeeRepositoryMongo {
	employeeRepo:= &EmployeeRepositoryMongo{}
	return employeeRepo;
}

func (erm EmployeeRepositoryMongo) CreateEmployee(employee *modal.Employee) error {
	db := database.GetDatabaseManager()
	if err := db.Create("employees", employee); err != nil {
		return err
	}

	return nil
}

func (erm EmployeeRepositoryMongo) GetEmployees() ([]modal.Employee, error) {
	db := database.GetDatabaseManager()

	result := []modal.Employee{}
	query := &bson.M{}
	if err := db.GetAllByQuery("employees", query, &result); err != nil {
		return result, err
	}

	return result, nil
}

func (erm EmployeeRepositoryMongo) UpdateEmployee(employee *modal.Employee) error {
	db := database.GetDatabaseManager()
	if err := db.Save("employees", employee.Id, employee); err != nil {
		return err
	}

	return nil
}