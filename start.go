
package main

import (
  "net/http"
  "fmt"
)

func (s *Server) Start() error {
  http.HandleFunc("/", HttpHandler(s))
  s.Log(fmt.Sprintf(M_server_start, s.Config.Listen))
  err := http.ListenAndServe(s.Config.Listen, nil)
  return err
}

