
package main

import (
  "fmt"
)

func (s Server) PrintLog(level uint, str string) {
  if level < s.Config.LogLevel  {
    return
  }
  fmt.Println(str)
}

func (s Server) Log(str string) {
  if log_level_log < s.Config.LogLevel {
    return
  }
  fmt.Println(str)
}

func (s Server) Warn(str string) {
  if log_level_warn < s.Config.LogLevel {
    return
  }
  fmt.Println(str)
}

func (s Server) Err(str string) {
  if log_level_error < s.Config.LogLevel {
    return
  }
  fmt.Println(str)
}

