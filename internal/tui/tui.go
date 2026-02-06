package tui

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/meklasdev/thunder-tui/internal/request"
	"gopkg.in/yaml.v3"
)

// Styles
var (
	// Colors
	primaryColor  = lipgloss.Color("#F7B93E")
	successColor  = lipgloss.Color("#7AA2F7")
	errorColor    = lipgloss.Color("#FF5555")
	bgColor       = lipgloss.Color("#1a1b26")
	borderColor   = lipgloss.Color("#7AA2F7")
	selectedColor = lipgloss.Color("#7dcfff")
	dimColor      = lipgloss.Color("#565f89")

	// Styles
	titleStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true).
			Padding(0, 1)

	selectedItemStyle = lipgloss.NewStyle().
				Foreground(selectedColor).
				Bold(true).
				PaddingLeft(2)

	itemStyle = lipgloss.NewStyle().
			Foreground(dimColor).
			PaddingLeft(2)

	panelStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(borderColor).
			Padding(1, 2)

	responsePanelStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(successColor).
				Padding(1, 2)

	statusBarStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Background(bgColor).
			Padding(0, 1)

	helpStyle = lipgloss.NewStyle().
			Foreground(dimColor).
			Padding(0, 1)
)

type model struct {
	collection     request.Collection
	selectedIndex  int
	response       string
	loading        bool
	err            error
	width          int
	height         int
	collectionPath string
}

type responseMsg struct {
	response request.Response
}

func NewModel(collectionPath string) model {
	// Load collection
	data, err := os.ReadFile(collectionPath)
	if err != nil {
		return model{
			err:            fmt.Errorf("failed to read collection: %w", err),
			collectionPath: collectionPath,
		}
	}

	var collection request.Collection
	if err := yaml.Unmarshal(data, &collection); err != nil {
		return model{
			err:            fmt.Errorf("failed to parse collection: %w", err),
			collectionPath: collectionPath,
		}
	}

	if len(collection.Requests) == 0 {
		return model{
			err:            fmt.Errorf("no requests found in collection"),
			collectionPath: collectionPath,
		}
	}

	return model{
		collection:     collection,
		selectedIndex:  0,
		response:       "Press Enter to send a request",
		collectionPath: collectionPath,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "up", "k":
			if m.selectedIndex > 0 {
				m.selectedIndex--
			}

		case "down", "j":
			if m.selectedIndex < len(m.collection.Requests)-1 {
				m.selectedIndex++
			}

		case "enter":
			if !m.loading && len(m.collection.Requests) > 0 {
				m.loading = true
				m.response = "⏳ Sending request..."
				return m, m.sendRequest()
			}
		}

	case responseMsg:
		m.loading = false
		m.response = request.FormatResponse(msg.response)

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("\n❌ Error: %v\n\nPress q to quit\n", m.err)
	}

	// Title
	title := titleStyle.Render("⚡ Thunder-TUI")

	// Request list
	var requestList strings.Builder
	requestList.WriteString(lipgloss.NewStyle().Foreground(primaryColor).Bold(true).Render("📋 Requests") + "\n\n")

	for i, req := range m.collection.Requests {
		methodColor := getMethodColor(req.Method)
		method := lipgloss.NewStyle().Foreground(methodColor).Bold(true).Render(fmt.Sprintf("%-6s", req.Method))

		if i == m.selectedIndex {
			line := selectedItemStyle.Render(fmt.Sprintf("▶ %s %s", method, req.Name))
			requestList.WriteString(line + "\n")
		} else {
			line := itemStyle.Render(fmt.Sprintf("  %s %s", method, req.Name))
			requestList.WriteString(line + "\n")
		}
	}

	leftPanel := panelStyle.Width(40).Height(m.height - 8).Render(requestList.String())

	// Response panel
	responseTitle := lipgloss.NewStyle().Foreground(successColor).Bold(true).Render("📡 Response")
	responseContent := responseTitle + "\n\n" + m.response
	rightPanel := responsePanelStyle.Width(m.width - 50).Height(m.height - 8).Render(responseContent)

	// Combine panels
	panels := lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, rightPanel)

	// Status bar
	selectedReq := m.collection.Requests[m.selectedIndex]
	status := statusBarStyle.Render(fmt.Sprintf("📂 %s | 🎯 %s %s", m.collectionPath, selectedReq.Method, selectedReq.URL))

	// Help
	help := helpStyle.Render("↑/↓: navigate | Enter: send | q: quit")

	// Combine all
	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		panels,
		"",
		status,
		help,
	)
}

func (m model) sendRequest() tea.Cmd {
	return func() tea.Msg {
		req := m.collection.Requests[m.selectedIndex]
		resp := request.Send(req)
		return responseMsg{response: resp}
	}
}

func getMethodColor(method string) lipgloss.Color {
	switch method {
	case "GET":
		return lipgloss.Color("#7AA2F7")
	case "POST":
		return lipgloss.Color("#9ECE6A")
	case "PUT":
		return lipgloss.Color("#F7B93E")
	case "DELETE":
		return lipgloss.Color("#FF5555")
	case "PATCH":
		return lipgloss.Color("#BB9AF7")
	default:
		return lipgloss.Color("#565f89")
	}
}
