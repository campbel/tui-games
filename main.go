package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	borderStyle = lipgloss.NewStyle().Background(lipgloss.Color("#ffffff")).Foreground(lipgloss.Color("#ffffff"))
)

type unit struct {
	char string
	pos  position
}

type position struct {
	x, y int
}

type window struct {
	width, height int
}

type model struct {
	booble   unit
	enemies  []unit
	window   window
	loading  bool
	quitting bool
}

func generateEnemies(xMax, yMax int) []unit {
	var enemies []unit
	for i := 0; i < 10; i++ {
		randX := rand.Intn(xMax-2) + 1
		randY := rand.Intn(yMax-2) + 1
		enemies = append(enemies, unit{
			char: "x",
			pos: position{
				x: randX,
				y: randY,
			},
		})
	}
	return enemies
}

func newModel() model {
	return model{
		booble: unit{
			char: "o",
			pos: position{
				x: 1,
				y: 1,
			},
		},
		enemies: []unit{},
		window: window{
			width:  0,
			height: 0,
		},
		loading: true,
	}
}

type tickMsg struct{}

func tickerCmd() tea.Cmd {
	timer := time.NewTicker(1 * time.Second)
	return func() tea.Msg {
		for range timer.C {
			return tickMsg{}
		}
		return nil
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
			m.booble.pos.y--
			if m.booble.pos.y < 1 {
				m.booble.pos.y = 1
			}
		case "a":
			m.booble.pos.x -= 2
			if m.booble.pos.x < 1 {
				m.booble.pos.x = 1
			}
		case "s":
			m.booble.pos.y++
			if m.booble.pos.y > m.window.height-2 {
				m.booble.pos.y = m.window.height - 2
			}
		case "d":
			m.booble.pos.x += 2
			if m.booble.pos.x > m.window.width-2 {
				m.booble.pos.x = m.window.width - 2
			}
		}
	case tickMsg:
		// Do something here
	case tea.WindowSizeMsg:
		m.window.width = msg.Width
		m.window.height = msg.Height - 1
		if m.loading {
			m.enemies = generateEnemies(m.window.width, m.window.height)
			m.loading = false
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.loading {
		return "Loading..."
	}

	s := ""

	for y := 0; y < m.window.height; y++ {
	XLOOP:
		for x := 0; x < m.window.width; x++ {
			if x == 0 || x == m.window.width-1 {
				s += borderStyle.Render(" ")
			} else if y == 0 || y == m.window.height-1 {
				s += borderStyle.Render(" ")
			} else if m.booble.pos.x == x && m.booble.pos.y == y {
				s += m.booble.char
			} else {
				for _, enemy := range m.enemies {
					if enemy.pos.x == x && enemy.pos.y == y {
						s += enemy.char
						continue XLOOP
					}
				}
				s += " "
			}
		}
		s += "\n"
	}

	return s
}

func main() {
	if _, err := tea.NewProgram(newModel(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
