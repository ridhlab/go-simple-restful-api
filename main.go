package main

import (
	"github/com/ridhlab/go-simple-restful-api/cmd"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	cmd.Execute()
}
