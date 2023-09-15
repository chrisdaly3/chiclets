package tui

import (
	"github.com/charmbracelet/bubbles/key"
)

/* KEYMAP */

type keys struct {
	Select    key.Binding
	Up        key.Binding
	Down      key.Binding
	Left      key.Binding
	Right     key.Binding
	Search    key.Binding
	Backspace key.Binding
	Back      key.Binding
	Quit      key.Binding
}

var Keybindings = keys{
	Select: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	),
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("▲/k", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("▼/j", "down"),
	),
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("◀/h", "left"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("▶/l", "right"),
	),
	Search: key.NewBinding(
		key.WithKeys("CAPSLOCK"),
		key.WithHelp("CAPS", "search"),
	),
	Backspace: key.NewBinding(
		key.WithKeys("backspace"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "go back"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("q/ctrl+c", "quit!"),
	),
}
