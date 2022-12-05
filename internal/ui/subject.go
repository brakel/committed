package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

const (
	subjectLimit int = 50
)

func subjectBlock(m model) string {
	return lipgloss.NewStyle().
		MarginBottom(1).
		Render(subjectRow(m.config.Emoji, m.config.Summary))
}

func subjectRow(e, s string) string {
	i := len(s)
	if e != "" {
		i += 2
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		emoji(e),
		summary(s),
		counter(i, subjectLimit),
	)
}

func emoji(str string) string {
	return lipgloss.NewStyle().
		Width(4).
		Height(1).
		MarginLeft(4).
		MarginRight(1).
		Align(lipgloss.Center, lipgloss.Center).
		BorderStyle(lipgloss.NormalBorder()).
		Render(str)
}

func summary(str string) string {
	return lipgloss.NewStyle().
		Width(61).
		Height(1).
		MarginRight(1).
		Align(lipgloss.Left, lipgloss.Center).
		Padding(0, 0, 0, 1).
		BorderStyle(lipgloss.NormalBorder()).
		Faint(true).
		Render(str)
}

func counter(count, total int) string {
	c := colour(fmt.Sprintf("%d", count), white)
	t := colour(fmt.Sprintf("%d", total), white)

	return lipgloss.NewStyle().
		Width(5).
		Height(3).
		Align(lipgloss.Right, lipgloss.Center).
		Render(fmt.Sprintf("%s/%s", c, t))
}
