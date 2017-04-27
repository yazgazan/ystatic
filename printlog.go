package main

import (
	"fmt"
)

func (s Server) PrintLog(level uint, str string) {
	if level < s.Config.LogLevel {
		return
	}
	fmt.Println(str)
}

func (s Server) Log(str string) {
	if logLevelLog < s.Config.LogLevel {
		return
	}
	fmt.Println(str)
}

func (s Server) Warn(str string) {
	if logLevelWarn < s.Config.LogLevel {
		return
	}
	fmt.Println(str)
}

func (s Server) Err(str string) {
	if logLevelError < s.Config.LogLevel {
		return
	}
	fmt.Println(str)
}
