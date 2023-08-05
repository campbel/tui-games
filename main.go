package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
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

type unit struct {
	width, height int
	char          func(u unit) string
	pos           position
	vector        vector
	stable        bool
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
	level    level
	window   window
	loading  bool
	quitting bool
}

type level struct {
	content []string
	width   int
	height  int
}

func newLevel(content []string) level {
	width := 0
	height := len(content)
	for _, line := range content {
		if len(line) > width {
			width = len(line)
		}
	}
	return level{
		content: content,
		width:   width,
		height:  height,
	}
}

func newModel(level []string) model {
	return model{
		booble: unit{
			height: 4,
			width:  5,
			char: func(u unit) string {
				var s string
				switch u.vector.x {
				case 1, 2:
					s = `
  O/
 /|
 / \
<>-<>`
				case -1, -2:
					s = `
 \O
  |\
 / \
<>-<>`
				default:
					s = `
  O
 /|\
 / \
<>-<>`
				}
				return strings.TrimLeft(s, "\n")
			},
			pos: position{
				x: 0,
				y: 0,
			},
			vector: vector{
				x: 0,
				y: 0,
			},
			stable: false,
		},
		level: newLevel(level),
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
		case "w", "W", " ", "up":
			if m.booble.stable {
				m.booble.vector.y = -4
			}
		case "a", "A", "left":
			if m.booble.stable {
				if m.booble.vector.x > 0 {
					m.booble.vector.x = 0
				} else if m.booble.vector.x > -2 {
					m.booble.vector.x--
				}
			}
		case "s", "S":
		case "d", "D", "right":
			if m.booble.stable {
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
		if m.booble.vector.y < 0 {
			m.booble.pos.y += m.booble.vector.y
			m.booble.vector.y++
		}
		for i := 0; i < m.booble.vector.y; i++ {
			if m.booble.pos.y+m.booble.height > m.level.height {
				return m, tea.Quit
			}
			for j := 0; j < m.booble.width; j++ {
				if m.level.content[m.booble.pos.y+m.booble.height+1][m.booble.pos.x+j] == '-' {
					m.booble.vector.y = 0
					break
				}
			}
			m.booble.pos.y++
		}
		m.booble.stable = false
		for j := 0; j < m.booble.width; j++ {
			if m.level.content[m.booble.pos.y+m.booble.height][m.booble.pos.x+j] == '-' {
				m.booble.stable = true
				break
			}
		}
		if !m.booble.stable {
			m.booble.vector.y++
		}

		// x component
		m.booble.pos.x += m.booble.vector.x
		if m.booble.pos.x < 0 {
			m.booble.pos.x = 0
			m.booble.vector.x = 0
		}
		if m.booble.pos.x > m.level.width-m.booble.width {
			m.booble.pos.x = m.level.width - m.booble.width
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
	if m.level.height+5 > m.window.height {
		return "Window is to small, please resize."
	}

	s := drawLevel(m.window, m.level)
	s = drawUnit(s, m.window, m.level, m.booble)
	s = drawInfo(s, m.window, m.level, m.booble)
	s = drawFrame(s, m.window, m.level, m.booble)

	return strings.Join(s, "\n")
}

const (
	minHeight = 20
	minWidth  = 80
)

func drawLevel(window window, level level) []string {
	if window.width < minWidth || window.height < level.height {
		return []string{"Terminal window is too small. Please resize and try again."}
	}
	offset := window.height - level.height
	result := make([]string, window.height)
	for i := 0; i < window.height; i++ {
		if i < offset {
			result[i] = stringRepeat(" ", window.width)
			continue
		}
		result[i] = level.content[i-offset] + stringRepeat(" ", window.width-len(level.content[i-offset]))
	}
	return result
}

func drawUnit(s []string, window window, level level, u unit) []string {
	lines := strings.Split(u.char(u), "\n")
	offset := window.height - level.height
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		posy := offset + u.pos.y + i
		s[posy] = s[posy][:u.pos.x] + line + s[posy][u.pos.x+len(line):]
	}
	return s
}

func drawInfo(s []string, window window, level level, u unit) []string {
	s[0] = fmt.Sprintf("X: %d Y: %d dX: %d, dY: %d", u.pos.x, u.pos.y, u.vector.x, u.vector.y)
	return s
}

func drawFrame(s []string, window window, level level, u unit) []string {
	// draw the relevant frame in the window
	var result = make([]string, len(s))

	for i := len(s) - 1; i >= 0; i-- {
		if i < window.height-level.height {
			result[i] = s[i]
			continue
		}

		startFrame := 0
		if u.pos.x > window.width/2 {
			startFrame = u.pos.x - window.width/2
		}
		if startFrame+window.width > level.width {
			startFrame = level.width - window.width
		}
		result[i] = s[i][startFrame : startFrame+window.width]
	}
	return result
}

func stringRepeat(s string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}

func main() {
	data, err := os.ReadFile("level.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	lines := processFile(data)

	if _, err := tea.NewProgram(newModel(lines), tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func processFile(data []byte) []string {
	lines := strings.Split(string(data), "\n")
	width := 0
	for _, line := range lines {
		if len(line) > width {
			width = len(line)
		}
	}
	for i, line := range lines {
		lines[i] = line + stringRepeat(" ", width-len(line))
	}
	return lines
}
