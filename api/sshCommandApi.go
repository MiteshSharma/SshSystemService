package api

import (
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/MiteshSharma/SshSystemSetup/utils"
	"github.com/MiteshSharma/SshSystemSetup/service"
	"github.com/MiteshSharma/SshSystemSetup/modal"
)

type SshCommandsResponse struct  {
	Commands []modal.SSHCommand
	Error error
}

func NewSshCommandsResponse(commands []modal.SSHCommand, error error) SshCommandsResponse  {
	commandsResponse := SshCommandsResponse{}
	commandsResponse.Error = error
	commandsResponse.Commands = commands
	return commandsResponse
}

type SshCommandResponse struct  {
	Command modal.SSHCommand
	Error error
}

func NewSshCommandResponse(command modal.SSHCommand, error error) SshCommandResponse  {
	commandResponse := SshCommandResponse{}
	commandResponse.Error = error
	commandResponse.Command = command
	return commandResponse
}

func InitSshCommandApi(router *httprouter.Router) {
	sshCommandApi := SshCommandApi{service.NewSshCommandService()};
	router.POST("/sshCommand", sshCommandApi.create)
	router.PUT("/sshCommand", sshCommandApi.update)
	router.GET("/sshCommand", sshCommandApi.getAll)
}

type SshCommandApi struct  {
	sshCommandService *service.SshCommandService
}

func (sca SshCommandApi) getAll(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sshCommands, err := sca.sshCommandService.GetSshCommands();
	if (err != nil) {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(utils.ToJson(NewSshCommandsResponse(sshCommands, err))))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(utils.ToJson(NewSshCommandsResponse(sshCommands, nil))))
}

func (sca SshCommandApi) create(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var command modal.SSHCommand
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&command); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(utils.ToJson("Incorrect body received.")))
		return
	}

	if !command.IsValid() {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(utils.ToJson("Command is not valid.")))
		return
	}

	// Run command using ssh command service
	command, err := sca.sshCommandService.CreateSshCommand(command)

	if (err != nil) {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(utils.ToJson(NewSshCommandResponse(command, err))))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(utils.ToJson(NewSshCommandResponse(command, nil))))
}

func (sca SshCommandApi) update(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var command modal.SSHCommand
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&command); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(utils.ToJson("Incorrect body received.")))
		return
	}

	if !command.IsValid() {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(utils.ToJson("Command is not valid.")))
		return
	}

	if command.Id == "" {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(utils.ToJson("Command is not created already, Can't update.")))
		return
	}

	// Run command using ssh command service
	command, err := sca.sshCommandService.UpdateSshCommand(command)

	if (err != nil) {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(utils.ToJson(NewSshCommandResponse(command, err))))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(utils.ToJson(NewSshCommandResponse(command, nil))))
}