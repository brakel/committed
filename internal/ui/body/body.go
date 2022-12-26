package body

import (
	"strings"

	"github.com/acarl005/stripansi"
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mikelorant/committed/internal/commit"
)

type Model struct {
	Height int
	Width  int

	focus    bool
	styles   Styles
	textArea textarea.Model
}

const (
	tabSize = 4

	defaultWidth = 72
)

func New(cfg commit.Config, h int) Model {
	m := Model{
		Height:   h,
		styles:   defaultStyles(),
		textArea: newTextArea(cfg.Placeholders.Body, defaultWidth),
	}

	if cfg.Amend {
		m.textArea.SetValue(cfg.HeadCommit.Body)
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

//nolint:ireturn
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd

	if m.focus {
		//nolint:gocritic
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "tab":
				if m.textArea.Focused() {
					m.textArea.InsertString(strings.Repeat(" ", tabSize))
				}
			}
		}
	}

	m.textArea.SetHeight(m.Height)

	switch {
	case m.focus && !m.textArea.Focused():
		cmd = m.textArea.Focus()
		return m, cmd
	case !m.focus && m.textArea.Focused():
		m.textArea.Blur()
		return m, nil
	}

	m.textArea, cmd = m.textArea.Update(msg)

	return m, cmd
}

func (m Model) View() string {
	return m.styles.boundary.Height(m.Height).Render(m.textArea.View())
}

func (m *Model) Focus() {
	m.focus = true
}

func (m *Model) Blur() {
	m.focus = false
}

func (m Model) Focused() bool {
	return m.focus
}

func (m Model) Value() string {
	if m.textArea.Value() == "" {
		return ""
	}

	m.textArea.Cursor.SetMode(cursor.CursorHide)
	res := strings.TrimSpace(stripansi.Strip(m.textArea.View()))
	m.textArea.Cursor.SetMode(cursor.CursorBlink)

	return res
}

func newTextArea(ph string, w int) textarea.Model {
	ta := textarea.New()
	ta.Placeholder = ph
	ta.Prompt = ""
	ta.ShowLineNumbers = false
	ta.SetWidth(w)

	return ta
}
