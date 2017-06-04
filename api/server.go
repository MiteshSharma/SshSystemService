package api

import (
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
	"github.com/MiteshSharma/SshSystemSetup/utils"
	"github.com/MiteshSharma/SshSystemSetup/middleware"
)

type Server struct  {
	Router *httprouter.Router
}

var ServerObj *Server

func InitServer()  {
	ServerObj = &Server{}
	ServerObj.Router = InitApi()
}

func StartServer()  {
	go func() {
		negroni := negroni.Classic()
		negroni.Use(middleware.NewRequest())
		negroni.Use(middleware.NewDebug())
		negroni.UseHandler(ServerObj.Router)
		negroni.Run(utils.ConfigParam.ServerConfig.Port)
	}()
}

func StopServer()  {
	// Closing DB connection
}