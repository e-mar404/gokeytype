package colors

import "github.com/charmbracelet/lipgloss"

type color string 
const (
	GRAY = "#bdbdbd"
  WHITE = "#ecf1ec"
  LAVENDER = "#89b4fa"
  PINK = "#f5c2e7"
  RED = "#f38ba8"
  TEXT = "#cdd6f4"
  BASE = "#1e1e2e"
  CRUST = "#11111b"
  SURFACE_0 = "#313244"
  OVERLAY_1 = "#7f849c"
)

var (
  AppStyle = lipgloss.NewStyle(). 
    Foreground(lipgloss.Color(TEXT)). 
    Background(lipgloss.Color(BASE)).
    Align(lipgloss.Center)
)

func Foreground(color color) lipgloss.Style {
  return lipgloss.NewStyle().Foreground(lipgloss.Color(color))
}

func Background(color color) lipgloss.Style {
  return lipgloss.NewStyle().Background(lipgloss.Color(color))
}
