package api

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/MiteshSharma/SshSystemSetup/utils"
	"github.com/MiteshSharma/SshSystemSetup/service"
	"github.com/MiteshSharma/SshSystemSetup/modal"
	"github.com/MiteshSharma/SshSystemSetup/repository"
)

type InstanceDetailResponse struct  {
	InstanceDetails []modal.InstanceDetail
	Error error
}

func NewInstanceDetailResponse(instances []modal.InstanceDetail, err error) InstanceDetailResponse  {
	taskResponse := InstanceDetailResponse{}
	taskResponse.InstanceDetails = instances
	taskResponse.Error = err
	return taskResponse
}

func InitInstanceApi(router *httprouter.Router) {
	instanceDetailApi := InstanceDetailApi{};
	router.GET("/instance", instanceDetailApi.getAll)
}

type InstanceDetailApi struct  {
}

func (id InstanceDetailApi) getAll(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ec2Service := service.NewInstanceService(repository.GetInstanceRepository("aws"));
	instances, err := ec2Service.GetDetails();
	if (err != nil) {
		rw.WriteHeader(http.StatusInternalServerError)
	} else {
		rw.WriteHeader(http.StatusOK)
	}
	rw.Write([]byte(utils.ToJson(NewInstanceDetailResponse(instances, err))))
}