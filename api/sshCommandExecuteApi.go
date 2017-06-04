package api

import (
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/MiteshSharma/SshSystemSetup/service"
	"github.com/MiteshSharma/SshSystemSetup/utils"
)

type SshCommandExecuteResponse struct  {
	IsSuccess bool
	Error error
}

func NewSshCommandExecuteResponse(error error) SshCommandResponse  {
	commandResponse := SshCommandExecuteResponse{}
	commandResponse.Error = error
	if (error != nil) {
		commandResponse.IsSuccess = false
	} else {
		commandResponse.IsSuccess = true
	}
	return commandResponse
}

func InitSshCommandExecuteApi(router *httprouter.Router) {
	sshCommandApi := SshCommandExecuteApi{service.NewSshCommandService()};
	router.POST("/sshCommandExecute", sshCommandApi.create)
}

type SshCommandExecuteApi struct  {
	sshCommandService *service.SshCommandService
}

func (sca SshCommandExecuteApi) create(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var commandObj service.SshCommandObj
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&commandObj); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(utils.ToJson("Incorrect body received.")))
		return
	}

	if !commandObj.Command.IsValid() {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(utils.ToJson("Command is not valid.")))
		return
	}

	if !commandObj.Instance.IsValid() {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(utils.ToJson("Instance details is not valid.")))
		return
	}

	// Run command using ssh command service
	err := sca.sshCommandService.RunCommand(commandObj)

	if (err != nil) {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(utils.ToJson(NewSshCommandExecuteResponse(err))))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(utils.ToJson(NewSshCommandExecuteResponse(nil))))
}
