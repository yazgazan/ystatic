
package main

const (
  log_level_log   = iota
  log_level_warn  = iota
  log_level_error = iota
)

const C_serverConfigFilename = "config.json"
const C_serverVarsFilename = "vars.json"

const C_configIndexesMax = 16
const C_configTemplatesMax = 16
const C_configMarkdownMax = 16
const C_configDefaultDelimiterO = "{{"
const C_configDefaultDelimiterC = "}}"

const C_defaultListen = "localhost:1234"
const C_defaultLogLevel = log_level_log
const C_defaultRoot = "www"

