package main

import (
	"github.com/caffeines/choto/cmd"
	"github.com/caffeines/choto/server"
)

func main() {
	cmd.Execute()
	server.RunServer()
}
