package main

import "net/http"

func Handler(server *Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := server.Router(w, r)
		if err != nil {
			http.Error(w, err.Error(), err.Code)
		}
	}
}
