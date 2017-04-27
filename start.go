package main

import (
	"fmt"
	"net/http"
)

func (s *Server) Start() error {
	http.HandleFunc("/", Handler(s))
	s.Log(fmt.Sprintf(serverStart, s.Config.Listen))
	err := http.ListenAndServe(s.Config.Listen, nil)
	return err
}
