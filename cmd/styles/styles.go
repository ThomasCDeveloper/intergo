package styles

import "github.com/charmbracelet/lipgloss"

var appWidth = 80
var Styles = map[string]lipgloss.Style{}

func InitStyles() {
	base := lipgloss.NewStyle().Width(appWidth)

	Styles["base"] = base
	Styles["highlighted"] = base.Bold(true)
	Styles["border"] = lipgloss.NewStyle().Reverse(true).Bold(true)
}

func Render(s string, styleName string) string {
	if style, ok := Styles[styleName]; ok {
		return style.Render(s)
	} else {
		return ""
	}
}
