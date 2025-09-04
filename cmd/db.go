package cmd

import (
	"fmt"
	"github/com/ridhlab/go-simple-restful-api/pkg"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

var dbCmd = &cobra.Command{
	Use:   "db [param]",
	Short: "Database Command",
	Run: func(cmd *cobra.Command, args []string) {
		param := args[0]
		if param == "init-db" {
			connDb := pkg.InitDB(os.Getenv("DB_SERVER"))
			sqlBytes, err := os.ReadFile("schema.sql")
			if err != nil {
				log.Printf("error reading schema.sql: %v", err)
				return
			}
			defer connDb.Close()
			_, err = connDb.Exec(string(sqlBytes))
			if err != nil {
				log.Printf("error executing schema.sql: %v", err)
				return
			}
			fmt.Println("success init db")
		} else {
			fmt.Println("unknown param")
		}
	},
}

func init() {
	rootCmd.AddCommand(dbCmd)
}
