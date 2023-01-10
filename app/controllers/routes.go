package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// type Server struct {
// 	DB     *gorm.DB
// 	Router *mux.Router
// }

func (server *Server) InitializeRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", Home).Methods("GET")
	server.Router.HandleFunc("/products", server.Products).Methods("GET")

	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/public/", http.FileServer(staticFileDirectory))
	server.Router.PathPrefix("/public/").Handler(staticFileHandler).Methods("GET")
}
