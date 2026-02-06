package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/meklasdev/thunder-tui/internal/tui"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("⚡ Thunder-TUI - Blazing Fast HTTP Client")
		fmt.Println("\nUsage:")
		fmt.Println("  thunder-tui run <collection.yaml>")
		fmt.Println("  thunder-tui help")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "run":
		if len(os.Args) < 3 {
			fmt.Println("❌ Error: Please specify a collection file")
			fmt.Println("Usage: thunder-tui run <collection.yaml>")
			os.Exit(1)
		}

		collectionPath := os.Args[2]

		// Initialize the TUI model
		m := tui.NewModel(collectionPath)

		// Start the Bubble Tea program
		p := tea.NewProgram(m, tea.WithAltScreen())

		if _, err := p.Run(); err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			os.Exit(1)
		}

	case "help", "--help", "-h":
		printHelp()

	default:
		fmt.Printf("❌ Unknown command: %s\n", command)
		fmt.Println("Run 'thunder-tui help' for usage information")
		os.Exit(1)
	}
}

func printHelp() {
	help := `
⚡ Thunder-TUI - Blazing Fast HTTP Client
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

USAGE:
  thunder-tui run <collection.yaml>

KEYBINDINGS:
  ↑/↓       Navigate requests
  Enter     Send request
  Tab       Switch panels
  q         Quit

COLLECTION FORMAT (YAML):
  requests:
    - name: "Get Users"
      method: GET
      url: "https://api.example.com/users"
      headers:
        Authorization: "Bearer token"
    
    - name: "Create User"
      method: POST
      url: "https://api.example.com/users"
      headers:
        Content-Type: "application/json"
      body: |
        {
          "name": "John Doe",
          "email": "john@example.com"
        }

EXAMPLES:
  thunder-tui run api-collection.yaml
  thunder-tui help

Made with ⚡ by meklasdev
`
	fmt.Println(help)
}
