package test

import (
	tea "github.com/charmbracelet/bubbletea"
)

const (
	GRAY = "\033[90m"
	RED = "\033[31m"
	GREEN = "\033[32m"
  CURSOR_BG = "\033[100m"
	RESET = "\033[0m"
)

type finishMessage string
type status int
type Test struct {
	text []byte
	status []status
	position int
}

const (
  Empty status = iota
	INCORRECT
	CORRECT
)

func New() Test {
  testText := "this is a longer test"
  statusSlice := make([]status, len(testText))
  for i := range(statusSlice) {
    statusSlice[i] = Empty 
  }

	return Test {
		text: []byte(testText),
		status: statusSlice, 
		position: 0,
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
		}

    t.position++
    if t.position == len(t.text) {
      return t, finishTest
    }
  case finishMessage:
    return t, tea.Quit
	}
	return t, nil
}

func (t Test) View() string {
	str := ""
	for i, letter := range(t.text) {
		switch t.status[i] {
		case CORRECT:
			str += GREEN + string(letter) + RESET
		case INCORRECT:
      if letter == ' ' {
        letter = '_'
      }
			str += RED + string(letter) + RESET
		default:
			str += GRAY + string(letter) + RESET
		}
	}

  
	return str + "\n"
}

func finishTest() tea.Msg {
  return finishMessage("Test is finished")
}
