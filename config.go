package main

type Config struct {
	Listen     string
	LogLevel   uint
	Root       string
	Frame      string
	Indexes    []string
	Templates  []string
	Markdown   []string
	Delimiters []string
	Framable   []string
}

func ConfigInit() *Config {
	ret := &Config{
		defaultListen,
		logLevelLog,
		defaultRoot,
		defaultFrame,
		make([]string, configIndexesMax),
		make([]string, configTemplatesMax),
		make([]string, configMarkdownMax),
		make([]string, 2),
		make([]string, configFramableMax),
	}
	ret.Delimiters[0] = configDefaultDelimiterO
	ret.Delimiters[1] = configDefaultDelimiterC
	return ret
}
