package styles

import "github.com/charmbracelet/lipgloss"

var appWidth = 80
var styles = map[string]lipgloss.Style{}

func InitStyles() {
	base := lipgloss.NewStyle().Width(appWidth)

	styles["base"] = base
	styles["highlighted"] = base.Bold(true)
}

func Render(s string, styleName string) string {
	if style, ok := styles[styleName]; ok {
		return style.Render(s)
	} else {
		return ""
	}
}
