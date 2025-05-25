package models

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/e-mar404/gokeytype/internal/colors"
)

var (
  logoStyle = colors.Foreground(colors.LAVENDER). 
    PaddingBottom(3).
    Background(lipgloss.Color(colors.BASE))

  optionsStyle = colors.Foreground(colors.PINK). 
    Inherit(colors.AppStyle)
)
