package cmd

import (
	"github.com/caffeines/choto/server"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve starts http server",
	Run:   serve,
}

func serve(cmd *cobra.Command, args []string) {
	server.RunServer()
}
