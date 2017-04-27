package main

import (
	"net/http"
)

func (s Server) Router(w http.ResponseWriter, r *http.Request) *ServerError {
	if err := s.HandleStaticFiles(w, r); err == nil {
		return nil
	}
	return &ServerError{
		fileNotFound,
		404,
	}
}
