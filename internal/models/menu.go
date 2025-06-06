package models 

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/e-mar404/gokeytype/internal/colors"
)

const (
  MainMenu state = iota
  WordCountMenu
)

type state int

type menu struct {
  state state
  logo string
  options []option
  selected int
  windowWidth int
  windowHeight int
}

type option struct {
  label string
  cmd tea.Cmd
}

type testWordCount int 

func NewMenu() menu {
  return menu {
    logo: ` ____ ____ ____ ____ ____ ____ ____ ____ ____
||G |||o |||k |||e |||y |||t |||y |||p |||e ||
||__|||__|||__|||__|||__|||__|||__|||__|||__||
|/__\|/__\|/__\|/__\|/__\|/__\|/__\|/__\|/__\|`,
    options: []option{
      option {
        label: "New Test",
        cmd: nil,
      },
      option {
        label: "Quit",
        cmd: nil,
      },
    },
    selected: 0,
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
    case "j":
      m.selected++
      if m.selected > len(m.options) - 1 {
        m.selected = len(m.options) - 1
      }
      return m, nil

    case "k":
      m.selected--
      if m.selected < 1 {
        m.selected = 0
      }
      return m, nil

    case "enter":
      switch m.state {
      case MainMenu:
        m.options = []option{
          option {
            label: "10 words",
            cmd: newTestWC(10),
          },
          option {
            label: "25 words",
            cmd: newTestWC(25),
          },
          option {
            label: "50 words",
            cmd: newTestWC(50),
          },
          option {
            label: "100 words",
            cmd: newTestWC(100),
          },
          option {
            label: "Quit",
            cmd: tea.Quit,
          },
        }
        m.state = WordCountMenu
        return m, nil 

      case WordCountMenu:
        return m, m.options[m.selected].cmd
      }

    case "q":
      return m, tea.Quit
    }

  case testWordCount:
    return newTest(int(msg), m.windowWidth, m.windowHeight), nil
  }


  return m, nil
}

func (m menu) View() string {
  var buf strings.Builder
  prefix := "  "
  buf.WriteString(colors.LogoStyle.Render(m.logo) + "\n")
  for i, option := range(m.options) {
    if i == m.selected {
      prefix = "> "
    } else {
      prefix = "  "
    }

    buf.WriteString(colors.OptionsStyle.Render(prefix + option.label) + "\n")
  }

  menu := buf.String()
  menuStyle := lipgloss.NewStyle(). 
    Inherit(colors.AppStyle).
    Width(m.windowWidth). 
    Height(m.windowHeight).
    PaddingTop(m.windowHeight/2 - (lipgloss.Height(menu)/2))

  return menuStyle.Render(menu)
}

func newTestWC(wc int) func() tea.Msg {
  return func() tea.Msg {
    return testWordCount(wc)
  } 
}
