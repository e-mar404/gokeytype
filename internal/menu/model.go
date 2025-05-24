package menu

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/e-mar404/gokeytype/internal/test"
)

type menu struct {
  logo string
  options []string
  windowWidth int
  windowHeight int
}

func New() menu {
  return menu {
    logo: `
 ____ ____ ____ ____ ____ ____ ____ ____ ____ 
||G |||o |||k |||e |||y |||t |||y |||p |||e ||
||__|||__|||__|||__|||__|||__|||__|||__|||__||
|/__\|/__\|/__\|/__\|/__\|/__\|/__\|/__\|/__\|`,
    options: []string{"n New Test", "q Quit"},
  }
}

func (m menu) Init() tea.Cmd {
  return nil
}

func (m menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.WindowSizeMsg:
    m.windowHeight = msg.Height
    m.windowWidth = msg.Width
  case tea.KeyMsg:
    switch msg.String() {
    case "n":
      return test.New(), nil
    case "q":
      return m, tea.Quit
    }
  }

  return m, nil
}

func (m menu) View() string {


  logoStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#00ADD8"))
  optionStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#CE3262"))

  var buf strings.Builder
  buf.WriteString(logoStyle.Render(m.logo))
  buf.WriteString("\n\n\n")
  for _, option := range(m.options) {
    buf.WriteString(optionStyle.Render(option) + "\n")
  }

  menu := buf.String()

  overallStyle := lipgloss.NewStyle().
    Width(m.windowWidth). 
    Height(m.windowHeight).
    PaddingTop(m.windowHeight/2 - (lipgloss.Height(menu)/2)).
    Align(lipgloss.Center)


  return overallStyle.Render(menu)
}
