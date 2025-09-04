package cmd

import (
	serverPkg "github/com/ridhlab/go-simple-restful-api/internal/server"
	"github/com/ridhlab/go-simple-restful-api/pkg"
	"os"

	"github.com/spf13/cobra"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Start the HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		connDb := pkg.InitDB(os.Getenv("DB_SERVER"))
		server := serverPkg.NewServer("localhost", "3030", connDb)
		defer connDb.Close()
		server.RunServer()
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
