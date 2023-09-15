package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	Black  = lipgloss.Color("#000000")
	Navy   = lipgloss.Color("#14213D")
	Violet = lipgloss.Color("#615789")
	Orange = lipgloss.Color("#FCA311")
	Grey   = lipgloss.Color("#E5E5E5")
	White  = lipgloss.Color("#FFFFFF")
)
var (
	FlexStyleBackground = lipgloss.NewStyle().
				Align(lipgloss.Center, lipgloss.Center).
				Background(Grey)

	FlexStyleText = lipgloss.NewStyle().
			Align(lipgloss.Center, lipgloss.Center).
			Foreground(Black).
			Background(Violet).
			Bold(true)

	FlexStyleNavy = lipgloss.NewStyle().
			Align(lipgloss.Center, lipgloss.Center).
			Background(Navy)

	FlexStyleOrange = lipgloss.NewStyle().
			Align(lipgloss.Center, lipgloss.Center).
			Background(Orange)

	FlexStyleWhite = lipgloss.NewStyle().
			Align(lipgloss.Center, lipgloss.Center).
			Background(White)

	FlexStyleViolet = lipgloss.NewStyle().
			Align(lipgloss.Center, lipgloss.Center).
			Background(Violet)

	FlexStyleBlank = lipgloss.NewStyle()
)
