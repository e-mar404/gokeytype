package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Status int
const (
	INCORRECT Status = iota
	CORRECT
	PENDING
)

const (
	GRAY = "\033[90m"
	RED = "\033[31m"
	GREEN = "\033[32m"
	RESET = "\033[0m"
)

type test struct {
	text []byte
	status []Status
	position int
}

func main() {
	p := tea.NewProgram(initialTest())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func initialTest() test {
	return test {
		text: []byte{'h', 'e', 'l', 'l', 'o'},
		status: []Status{PENDING, PENDING, PENDING, PENDING, PENDING}, 
		position: 0,
	}
}

func (t test) Init() tea.Cmd{
	return nil
}

func (t test) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		input := msg.String()
		if input == "ctrl+c" {
			return t, tea.Quit
		}

		if input == string(t.text[t.position]) {
			t.status[t.position] = CORRECT
		} else {
			t.status[t.position] = INCORRECT
		}

		t.position++
	}
	return t, nil
}

func (t test) View() string {
	str := ""
	for i, letter := range(t.text) {
		switch t.status[i] {
		case CORRECT:
			str += GREEN + string(letter) + RESET
		case INCORRECT:
			str += RED + string(letter) + RESET
		default:
			str += GRAY + string(letter) + RESET
		}
	}
	return str
}
