package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Init all routes here
func InitApi() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", get)
	InitInstanceApi(router)
	InitSshCommandApi(router)
	InitSshCommandExecuteApi(router)
	InitEmployeeApi(router)
	InitUserApi(router)
	return router
}

func get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rw.Header().Set("Content-Type", "application/json")
	rw.Write([]byte("{'welcome': 'hey'}"))
}