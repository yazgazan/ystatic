
package main

import (
  "github.com/russross/blackfriday"

  "net/http"
  "io/ioutil"
  "os"
  "fmt"
)

func (s Server) ToMarkdown(content []byte) []byte {
  return blackfriday.MarkdownCommon(content)
}

func (s Server) HandleMarkdown(
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
    file.Close()
    html := s.ToMarkdown(fileContent)
    w.Header()["Content-Length"] = make([]string, 1)
    w.Header()["Content-Length"][0] = fmt.Sprintf("%d", len(html))
    w.Write(html)
  return nil
}


