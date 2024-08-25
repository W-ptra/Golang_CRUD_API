package main

import (
	"log"
	"net/http"
	"github.com/W-ptra/Golang_CRUD_API/Router"
)

type APIServer struct {
	addr string
}

func newAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) run() error {
	router := http.NewServeMux()
	router.HandleFunc("GET /api/student",router1.StudentGet)
	router.HandleFunc("GET /api/student/{id}",router1.StudentGetById)
	router.HandleFunc("POST /api/student",router1.StudentPost)
	router.HandleFunc("PUT /api/student/{id}",router1.StudentPut)
	router.HandleFunc("DELETE /api/student/{id}",router1.StudentDelete)

	server := http.Server{
		Addr: s.addr,
		Handler: router,
	}

	log.Printf("Server has started %s",s.addr)

	return server.ListenAndServe()
}