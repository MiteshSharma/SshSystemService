package api

import (
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/MiteshSharma/SshSystemSetup/utils"
	"github.com/MiteshSharma/SshSystemSetup/modal"
	"github.com/MiteshSharma/SshSystemSetup/service"
)

type UsersResponse struct  {
	Users []modal.User
	Error error
}

func NewUsersResponse(users []modal.User, err error) UsersResponse  {
	userResponse := UsersResponse{}
	userResponse.Users = users
	userResponse.Error = err
	return userResponse
}

type UserResponse struct  {
	User modal.User
	Error error
}

func NewUserResponse(user modal.User, err error) UserResponse  {
	userResponse := UserResponse{}
	userResponse.User = user
	userResponse.Error = err
	return userResponse
}

func InitUserApi(router *httprouter.Router) {
	userApi := UserApi{};
	userApi.userService = service.NewUserService()
	router.GET("/user", userApi.getAll)
	router.POST("/user", userApi.create)
	router.PUT("/user", userApi.update)
}

type UserApi struct  {
	userService service.UserService
}

func (ua UserApi) getAll(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	users, err := ua.userService.GetUsers();
	if (err != nil) {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(utils.ToJson(NewUsersResponse(users, err))))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(utils.ToJson(NewUsersResponse(users, nil))))
}

func (ua UserApi) create(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var user modal.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(utils.ToJson("Incorrect body received.")))
		return
	}

	if !user.IsValid() {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(utils.ToJson("User data is not valid.")))
		return
	}

	// Run command using ssh command service
	user, err := ua.userService.CreateUser(user)

	if (err != nil) {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(utils.ToJson(NewUserResponse(user, err))))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(utils.ToJson(NewUserResponse(user, nil))))
}

func (ua UserApi) update(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var user modal.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(utils.ToJson("Incorrect body received.")))
		return
	}

	if !user.IsValid() {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(utils.ToJson("User data is not valid.")))
		return
	}

	if user.Id == "" {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(utils.ToJson("User is not created already, Can't update.")))
		return
	}

	// Run command using ssh command service
	user, err := ua.userService.UpdateUser(user)

	if (err != nil) {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(utils.ToJson(NewUserResponse(user, err))))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(utils.ToJson(NewUserResponse(user, nil))))
}
