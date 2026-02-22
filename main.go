package main

import (
	"fmt"
	"os"

	"ntrms/internal/model"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() // .env is optional; fall through to env vars if missing
	apiKey := os.Getenv("TWELVEDATA_API_KEY")

	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "TWELVEDATA_API_KEY environment variable is required")

		os.Exit(1)
	}

	p := tea.NewProgram(model.New(apiKey))
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)

		os.Exit(1)
	}
}
