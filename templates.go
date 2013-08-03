
package main

import (
  "net/http"
  "io/ioutil"
  "os"
  "html/template"
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
    tpl.Execute(w, variables)
  return nil
}

