package app

import (
	"github.com/gorilla/mux"
	"github.com/im4mbukh4ri/go-toko/app/controllers"
)

func (server *Server) InitializeRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}
