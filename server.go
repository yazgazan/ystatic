package main

import (
	"encoding/json"
	"io/ioutil"
)

type Server struct {
	Config    *Config
	Variables map[string]string
}

func (s *Server) LoadConfig() error {
	buf, err := ioutil.ReadFile(serverConfigFilename)
	if err != nil {
		return err
	}

	return json.Unmarshal(buf, &s.Config)
}

func (s *Server) LoadVars2() (map[string]string, error) {
	var err error
	ret := make(map[string]string)

	buf, err := ioutil.ReadFile(serverVarsFilename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(buf, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *Server) LoadVars() error {
	buf, err := ioutil.ReadFile(serverVarsFilename)
	if err != nil {
		return err
	}

	return json.Unmarshal(buf, &s.Variables)
}

func ServerInit() *Server {
	return &Server{
		ConfigInit(),
		make(map[string]string),
	}
}
