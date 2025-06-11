package models

import (
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/e-mar404/gokeytype/internal/colors"
	"github.com/e-mar404/gokeytype/internal/stats"
)

type result struct {
  windowWidth int
  windowHeight int
  options []string
  stats stats.Stats 
}

func newResult(stats stats.Stats, width, height int) result {
  return result {
    windowWidth: width,
    windowHeight: height,
    options: []string{"n New Test", "q Quit"},
    stats: stats,
  }
}

func (r result) Init() tea.Cmd {
  return nil
}

func (r result) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.String() {
    case "n":
      return newTest(r.stats.WordCount, r.windowWidth, r.windowHeight), nil

    case "q":
      return r, tea.Quit
    }
  }

  return r, nil
}

func (r result) View() string {
  var buf strings.Builder

  acc := 100 - (float64(r.stats.IncorrectChars) / float64(r.stats.TotalChars) * 100)
  accStr := strconv.FormatFloat(acc, 'f', 2, 32)
  buf.WriteString("Accuracy: " + accStr + "%\n")

  for _, option := range(r.options) {
    buf.WriteString(colors.OptionsStyle.Render(option) + "\n")
  }

  menu := buf.String()
  menuStyle := lipgloss.NewStyle(). 
    Inherit(colors.AppStyle).
    Width(r.windowWidth). 
    Height(r.windowHeight).
    PaddingTop(r.windowHeight/2 - (lipgloss.Height(menu)/2))

  return menuStyle.Render(menu)
}
