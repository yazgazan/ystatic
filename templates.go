
package main

import (
  "net/http"
  "io/ioutil"
  "os"
  "html/template"
  "bytes"
  "fmt"
)

func (s Server) HandleTemplate(
  w http.ResponseWriter,
  r *http.Request,
  file *os.File) error {
    fileContent, err := ioutil.ReadFile(file.Name())
    if err != nil {
      return &ServerError{
        M_server_404,
        404,
      }
    }
    if s.MatchExtension(file.Name(), s.Config.Markdown) == true {
      buf := s.ToMarkdown(fileContent)
      fileContent = buf
    }
    file.Close()
    tpl := template.New("main")
    tpl.Delims(s.Config.Delimiters[0], s.Config.Delimiters[1])
    tpl.Parse(string(fileContent))
    variables, VarsErr := s.LoadVars2()
    if VarsErr != nil {
      return &ServerError{
        M_vars_not_found,
        500,
      }
    }
    framable := s.MatchExtension(file.Name(), s.Config.Framable)
    if len(s.Config.Frame) == 0 || ! framable {
      tpl.Execute(w, variables)
      return nil
    }
    var buf bytes.Buffer
    tpl.Execute(&buf, variables)
    frameFileName := fmt.Sprintf("%s/%s", s.Config.Root, s.Config.Frame)
    frameContent, frameErr := ioutil.ReadFile(frameFileName)
    if frameErr != nil {
      return &ServerError{
        fmt.Sprintf(M_frame_failed, s.Config.Frame),
        500,
      }
    }
    frame := template.New("frame")
    frame.Delims(s.Config.Delimiters[0], s.Config.Delimiters[1])
    frame.Parse(string(frameContent))
    frame.Execute(w, template.HTML(buf.String()))
  return nil
}

