
package main

import (
  "os"
  "fmt"
  "path"
  "net/http"
)

func (s Server) TryIndexes(dirPath string) (*os.File, error) {
  for _, indexName := range s.Config.Indexes {
    pathName := fmt.Sprintf("%s%s", dirPath, indexName)
    file, err := s.TryFile(pathName)
    if err == nil {
      return file, nil
    }
  }
  return nil, &ServerError{M_server_404, 404}
}

func (s Server) TryFile(filePath string) (*os.File, error) {
  fileInfo, StatErr := os.Stat(filePath)
  if StatErr != nil {
    return nil, StatErr
  }
  if fileInfo.IsDir() == true {
    return s.TryIndexes(filePath)
  }
  return os.Open(filePath)
}

func (s Server) BuildStaticPath(path string) string {
  return fmt.Sprintf("%s%s", s.Config.Root, path)
}

func (s Server) MatchExtension(filePath string, exts []string) bool {
  fileExt := path.Ext(filePath)
  for _, ext := range exts {
    if ext == fileExt {
      return true
    }
  }
  return false
}

func (s Server) HandleStaticFiles(w http.ResponseWriter, r *http.Request) error {
  var file *os.File
  var err error
  staticPath := s.BuildStaticPath(r.URL.Path)
  file, err = s.TryFile(staticPath)
  if err != nil {
    return err
  }
  if s.MatchExtension(file.Name(), s.Config.Templates) == true {
    return s.HandleTemplate(w, r, file)
  }
  if s.MatchExtension(file.Name(), s.Config.Markdown) == true {
    return s.HandleMarkdown(w, r, file)
  }
  http.ServeFile(w, r, file.Name())
  file.Close()
  return nil
}

