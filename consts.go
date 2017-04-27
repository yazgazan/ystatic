package main

const (
	logLevelLog   = iota
	logLevelWarn  = iota
	logLevelError = iota
)

const serverConfigFilename = "config.json"
const serverVarsFilename = "vars.json"

const configIndexesMax = 16
const configTemplatesMax = 16
const configMarkdownMax = 16
const configFramableMax = 16
const configDefaultDelimiterO = "{{"
const configDefaultDelimiterC = "}}"

const defaultListen = "localhost:1234"
const defaultRoot = "www"
const defaultFrame = ""
