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
	Previous  key.Binding
	Backspace key.Binding
	Esc       key.Binding
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
	Previous: key.NewBinding(
		key.WithKeys("ctrl+b"),
		key.WithHelp("ctrl+b", "Prior Screen"),
	),
	Backspace: key.NewBinding(
		key.WithKeys("backspace"),
	),
	Esc: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "remove filter"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("q/ctrl+c", "quit!"),
	),
}

// HelpText relies on the key.WithHelp key.
// String formatting as follows
// Up, Down, Left, Right,
// Select,
// Previous
// Search,
// Esc,
// Quit
var HelpText = `
  Movement: %s, %s, %s, %s
  Select: %s | Prior Screen: %s
  Filter in Column: %s | Remove Filter: %s
  Quit: %s
  `
