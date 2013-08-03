
package main

import (
  "io/ioutil"
  "encoding/json"
)

type Server struct {
  Config    *Config
  Variables map[string]string
}

func (s *Server) LoadConfig() error  {
  buf, errReadFile := ioutil.ReadFile(C_serverConfigFilename)
  if errReadFile != nil  {
    return errReadFile
  }
  errJson := json.Unmarshal(buf, &s.Config)
  if errJson != nil  {
    return errJson
  }
  return nil
}

func (s *Server) LoadVars() error  {
  buf, errReadFile := ioutil.ReadFile(C_serverVarsFilename)
  if errReadFile != nil  {
    return errReadFile
  }
  errJson := json.Unmarshal(buf, &s.Variables)
  if errJson != nil  {
    return errJson
  }
  return nil
}

func ServerInit() *Server {
  return &Server{
    ConfigInit(),
    make(map[string]string),
  }
}

