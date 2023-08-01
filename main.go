package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var logger func(string, ...interface{})

func init() {
	f, err := os.Create("debug.txt")
	if err != nil {
		panic(err)
	}

	logger = func(format string, a ...any) {
		fmt.Fprintf(f, format, a...)
		fmt.Fprintln(f)
	}
}

var (
	borderStyle = lipgloss.NewStyle().Background(lipgloss.Color("#ffffff")).Foreground(lipgloss.Color("#ffffff"))
)

type unit struct {
	char   string
	pos    position
	vector vector
}

type position struct {
	x, y int
}

type vector struct {
	x, y int
}

type window struct {
	width, height int
}

type model struct {
	booble   unit
	window   window
	loading  bool
	quitting bool
}

func newModel() model {
	return model{
		booble: unit{
			char: `
  0
 /|\
 / \
~~~~~
"   "`,
			pos: position{
				x: 0,
				y: 0,
			},
			vector: vector{
				x: 0,
				y: 0,
			},
		},
		window: window{
			width:  0,
			height: 0,
		},
		loading: true,
	}
}

type tickMsg struct{}

func tickerCmd() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(time.Second / 24)
		return tickMsg{}
	}
}

func (m model) Init() tea.Cmd {
	return tickerCmd()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.quitting = true
			return m, tea.Quit
		case "w":
			if m.booble.pos.y == 0 {
				m.booble.vector.y = 4
			}
		case "a":
			if m.booble.pos.y == 0 {
				if m.booble.vector.x > 0 {
					m.booble.vector.x = 0
				} else if m.booble.vector.x > -2 {
					m.booble.vector.x--
				}
			}
		case "s":
		case "d":
			if m.booble.pos.y == 0 {
				if m.booble.vector.x < 0 {
					m.booble.vector.x = 0
				} else if m.booble.vector.x < 2 {
					m.booble.vector.x++
				}
			}
		}
		return m, nil
	case tickMsg:
		// y component
		m.booble.pos.y += m.booble.vector.y
		m.booble.vector.y--
		if m.booble.pos.y <= 0 {
			m.booble.pos.y = 0
			m.booble.vector.y = 0
		}
		// x component
		m.booble.pos.x += m.booble.vector.x
		if m.booble.pos.x < 0 {
			m.booble.pos.x = 0
			m.booble.vector.x = 0
		}
		if m.booble.pos.x > m.window.width-6 {
			m.booble.pos.x = m.window.width - 6
			m.booble.vector.x = 0
		}
		return m, tickerCmd()
	case tea.WindowSizeMsg:
		m.window.width = msg.Width
		m.window.height = msg.Height
		if m.loading {
			m.loading = false
		}
		return m, nil
	}
	return m, nil
}

func (m model) View() string {
	if m.loading {
		return "Loading..."
	}

	s := drawBackground(m.window.width, m.window.height)
	s = drawUnit(s, m.booble)

	return strings.Join(s, "\n")
}

const (
	minHeight = 20
	minWidth  = 80
)

func drawBackground(width, height int) []string {
	if width < minWidth || height < minHeight {
		return []string{"Terminal window is too small. Please resize and try again."}
	}
	result := []string{}
	for y := 0; y < height; y++ {
		if y == height-2 {
			result = append(result, stringRepeat("-", width))
		} else if y == height-1 {
			result = append(result, stringRepeat("/", width))
		} else {
			result = append(result, stringRepeat(" ", width))
		}
	}
	return result
}

func drawUnit(s []string, u unit) []string {
	lines := strings.Split(u.char, "\n")
	for i := len(lines) - 1; i > 0; i-- {
		line := lines[i]
		posY := i + len(s) - 2 - len(lines) - u.pos.y
		s[posY] = s[posY][:u.pos.x] + line + s[posY][u.pos.x+len(line):]
	}
	return s
}

func stringRepeat(s string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}

func main() {
	if _, err := tea.NewProgram(newModel(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
