package models 

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/e-mar404/gokeytype/internal/colors"
)

type menu struct {
  logo string
  options []string
  windowWidth int
  windowHeight int
}

func NewMenu() menu {
  return menu {
    logo: ` ____ ____ ____ ____ ____ ____ ____ ____ ____
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
    m.windowWidth, m.windowHeight = msg.Width, msg.Height

  case tea.KeyMsg:
    switch msg.String() {
    case "n":
      return newTest(m.windowWidth, m.windowHeight), nil
    case "q":
      return m, tea.Quit
    }
  }

  return m, nil
}

func (m menu) View() string {
  var buf strings.Builder
  buf.WriteString(logoStyle.Render(m.logo) + "\n")
  for _, option := range(m.options) {
    buf.WriteString(optionsStyle.Render(option) + "\n")
  }

  menu := buf.String()
  menuStyle := lipgloss.NewStyle(). 
    Inherit(colors.AppStyle).
    Width(m.windowWidth). 
    Height(m.windowHeight).
    PaddingTop(m.windowHeight/2 - (lipgloss.Height(menu)/2))

  return menuStyle.Render(menu)
}
