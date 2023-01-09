package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/im4mbukh4ri/go-toko/app/controllers"
)

func (server *Server) InitializeRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")

	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/public/", http.FileServer(staticFileDirectory))
	server.Router.PathPrefix("/public/").Handler(staticFileHandler).Methods("GET")
}
