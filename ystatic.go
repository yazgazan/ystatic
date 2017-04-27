package main

func main() {
	server := ServerInit()
	if err := server.LoadConfig(); err != nil {
		panic(err)
	}
	if err := server.LoadVars(); err != nil {
		server.Err(varsNotFound)
		server.Err(err.Error())
	}
	if err := server.Start(); err != nil {
		panic(err)
	}
}
