package main

import (
	"os"

	"github.com/tlkamp/pomodoro/internal/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
