
package main

type Config struct  {
  Listen      string
  LogLevel    uint
  Root        string
  Indexes     []string
  Templates   []string
  Markdown    []string
  Delimiters  []string
}

func ConfigInit() *Config {
  ret := &Config{
    C_defaultListen,
    log_level_log,
    C_defaultRoot,
    make([]string, C_configIndexesMax),
    make([]string, C_configTemplatesMax),
    make([]string, C_configMarkdownMax),
    make([]string, 2),
  }
  ret.Delimiters[0] = C_configDefaultDelimiterO
  ret.Delimiters[1] = C_configDefaultDelimiterC
  return ret
}

