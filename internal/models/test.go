package models

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/e-mar404/gokeytype/internal/colors"
	"github.com/e-mar404/gokeytype/internal/stats"
	"github.com/e-mar404/gokeytype/internal/text"
)

type finishMsg string
type status int
type Test struct {
  errorMsg string
	text []byte
	status []status
	position int
  windowWidth int
  windowHeight int
  stats stats.Stats 
}

const (
  Empty status = iota
	INCORRECT
	CORRECT
)

func newTest(wordCount, width, height int) Test {
  testText, err := text.Generate(wordCount)
  if err != nil {
    return Test {
      errorMsg: err.Error(),
    }
  }

  statusSlice := make([]status, len(testText))
  for i := range(statusSlice) {
    statusSlice[i] = Empty 
  }

	return Test {
		text: []byte(testText),
		status: statusSlice, 
		position: 0,
    windowWidth: width,
    windowHeight: height,
    stats: stats.Stats {
      WordCount: wordCount,
      TotalChars: len(testText),
      IncorrectChars: 0,
    },
	}
}

func (t Test) Init() tea.Cmd{
	return nil
}

func (t Test) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		input := msg.String()
		if input == "ctrl+c" {
			return t, tea.Quit
		}

    if input == "backspace" {
      if t.position > 0 {
        t.position--
        t.status[t.position] = Empty 
      }
      return t, nil
    }

		if input == string(t.text[t.position]) {
			t.status[t.position] = CORRECT
		} else {
			t.status[t.position] = INCORRECT
      t.stats.IncorrectChars++
		}

    t.position++
    if t.position == len(t.text) {
      return t, finishTest
    }

  case finishMsg:
    return newResult(t.stats, t.windowWidth, t.windowHeight), nil 
	}

	return t, nil
}

func (t Test) View() string {
  if t.errorMsg != "" {
    return t.errorMsg
  }

  str := make([]string, len(t.text))

	for i, letter := range(t.text) {
		switch t.status[i] {
		case CORRECT:
      coloredLetter := colors.CorrectAttemptStyle.Render(string(letter))

			str[i] = coloredLetter
		case INCORRECT:
      if letter == ' ' {
        letter = '_'
      }
      coloredLetter := colors.IncorrectAttemptStyle.Render(string(letter))

			str[i] = coloredLetter
		default:
      var coloredLetter string
      if i == t.position {
        coloredLetter = colors.CursorStyle.Render(string(letter))
      } else {
        coloredLetter = colors.EmptyAttemptStyle.Render(string(letter))
      }

			str[i] = coloredLetter
		} 
	}

  text := strings.Join(str, "")
  testStyle := lipgloss.NewStyle(). 
    Inherit(colors.AppStyle).
    Width(t.windowWidth). 
    Height(t.windowHeight).
    PaddingTop(t.windowHeight/2 - (lipgloss.Height(text)/2))

	return testStyle.Render(text) + "\n"
}

func finishTest() tea.Msg {
  return finishMsg("Test is finished")
}
