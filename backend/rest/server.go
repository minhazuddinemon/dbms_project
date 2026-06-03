package rest

import (
	"dbms-project/database/db"
	"net/http"
)

type Server struct {
	Store  db.Store
	Router *http.ServeMux
}

func NewServer(store db.Store) *Server {
	server := &Server{
		Store:  store,
		Router: http.NewServeMux(),
	}

	server.RegisterRoutes()
	return server
}

func (server *Server) RegisterRoutes() {
	server.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello, building a dbms projecttesting the docker "))
	})
}
