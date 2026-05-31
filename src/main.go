package main

import (
	"log/slog"
	"trains/src/cli/cmd"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Warn("Error loading .env file")
	}
	cmd.Execute()
}
