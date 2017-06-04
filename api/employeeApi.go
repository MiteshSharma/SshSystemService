package api

import (
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/MiteshSharma/SshSystemSetup/modal"
	"github.com/MiteshSharma/SshSystemSetup/service"
	"github.com/MiteshSharma/SshSystemSetup/utils"

)

type EmployeesResponse struct  {
	Employees []modal.Employee
	Error error
}

func NewEmployeesResponse(employees []modal.Employee, err error) InstanceDetailResponse  {
	employeesResponse := EmployeesResponse{}
	employeesResponse.Employees = employees
	employeesResponse.Error = err
	return employeesResponse
}

type EmployeeResponse struct  {
	Employee modal.Employee
	Error error
}

func NewEmployeeResponse(employee modal.Employee, err error) InstanceDetailResponse  {
	employeeResponse := EmployeeResponse{}
	employeeResponse.Employee = employee
	employeeResponse.Error = err
	return employeeResponse
}

func InitEmployeeApi(router *httprouter.Router) {
	employeeApi := EmployeeApi{};
	employeeApi.employeeService = service.NewEmployeeService()
	router.GET("/employee", employeeApi.getAll)
	router.POST("/employee", employeeApi.create)
	router.PUT("/employee", employeeApi.update)
}

type EmployeeApi struct  {
	employeeService service.EmployeeService
}

func (ea EmployeeApi) getAll(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	employees, err := ea.employeeService.GetEmployees();
	if (err != nil) {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(utils.ToJson(NewEmployeesResponse(employees, err))))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(utils.ToJson(NewEmployeesResponse(employees, nil))))
}

func (ea EmployeeApi) create(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var employee modal.Employee
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&employee); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(utils.ToJson("Incorrect body received.")))
		return
	}

	if !employee.IsValid() {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(utils.ToJson("Employee data is not valid.")))
		return
	}

	// Run command using ssh command service
	employee, err := ea.employeeService.CreateEmployee(employee)

	if (err != nil) {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(utils.ToJson(NewEmployeeResponse(employee, err))))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(utils.ToJson(NewEmployeeResponse(employee, nil))))
}

func (ea EmployeeApi) update(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var employee modal.Employee
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&employee); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(utils.ToJson("Incorrect body received.")))
		return
	}

	if !employee.IsValid() {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(utils.ToJson("Employee data is not valid.")))
		return
	}

	if employee.Id == "" {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(utils.ToJson("Employee is not created already, Can't update.")))
		return
	}

	// Run command using ssh command service
	employee, err := ea.employeeService.UpdateEmployee(employee)

	if (err != nil) {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(utils.ToJson(NewEmployeeResponse(employee, err))))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(utils.ToJson(NewEmployeeResponse(employee, nil))))
}
