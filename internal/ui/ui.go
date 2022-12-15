package ui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mikelorant/committed/internal/commit"
	"github.com/mikelorant/committed/internal/ui/body"
	"github.com/mikelorant/committed/internal/ui/footer"
	"github.com/mikelorant/committed/internal/ui/header"
	"github.com/mikelorant/committed/internal/ui/info"
	"github.com/mikelorant/committed/internal/ui/status"
)

type Model struct {
	state  state
	models Models
	result Result
	err    error
}

type Models struct {
	info   info.Model
	header header.Model
	body   body.Model
	footer footer.Model
	status status.Model
}

type Result struct {
	Commit  bool
	Name    string
	Email   string
	Emoji   string
	Summary string
	Body    string
	Footer  string
}

type state int

const (
	emptyComponent state = iota
	authorComponent
	emojiComponent
	summaryComponent
	bodyComponent
)

func New(cfg commit.Config) (Result, error) {
	logfilePath := os.Getenv("BUBBLETEA_LOG")
	if logfilePath != "" {
		fh, err := tea.LogToFile(logfilePath, "committed")
		if err != nil {
			return Result{}, fmt.Errorf("unable to log to file: %w", err)
		}
		defer fh.Close()
	}

	im := Model{
		state: emojiComponent,
		models: Models{
			info:   info.New(cfg),
			header: header.New(cfg),
			body:   body.New(cfg),
			footer: footer.New(cfg),
			status: status.New(),
		},
	}

	p := tea.NewProgram(im)
	m, err := p.Run()
	if err != nil {
		return Result{}, fmt.Errorf("unable to run program: %w", err)
	}

	return m.(Model).result, nil
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.models.info.Init(),
		m.models.header.Init(),
		m.models.body.Init(),
		m.models.footer.Init(),
		m.models.status.Init(),
	)
}

//nolint:ireturn
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//nolint:gocritic
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "alt+1":
			if m.state == authorComponent {
				return m, nil
			}
			m.state = authorComponent
		case "alt+2":
			if m.state == emojiComponent {
				return m, nil
			}
			m.state = emojiComponent
		case "alt+3":
			if m.state == summaryComponent {
				return m, nil
			}
			m.state = summaryComponent
		case "alt+4":
			if m.state == bodyComponent {
				return m, nil
			}
			m.state = bodyComponent
		case "enter":
			if m.state == emojiComponent {
				m.models.header, _ = m.models.header.Update(msg)
				m.state = summaryComponent
				break
			}
			if m.state == summaryComponent {
				m.state = bodyComponent
			}
		case "alt+enter":
			m.result = Result{
				Commit:  true,
				Name:    m.models.info.Name,
				Email:   m.models.info.Email,
				Emoji:   m.models.header.Emoji.ShortCode,
				Summary: m.models.header.Summary(),
				Body:    m.models.body.Value(),
			}
			if m.validate() {
				return m, tea.Quit
			}
		case "tab":
			switch m.state {
			case authorComponent:
				m.state = emojiComponent
			case emojiComponent:
				m.state = summaryComponent
			case summaryComponent:
				m.state = bodyComponent
			}
		case "shift+tab":
			switch m.state {
			case emojiComponent:
				m.state = authorComponent
			case summaryComponent:
				m.state = emojiComponent
			case bodyComponent:
				m.state = summaryComponent
			}
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	m.models.header.Blur()
	m.models.header.Expand = false
	m.models.body.Blur()
	m.models.body.Compact = false

	switch m.state {
	case emojiComponent:
		m.models.header.Focus()
		m.models.header.SelectEmoji()
		m.models.header.Expand = true
		m.models.body.Compact = true
	case summaryComponent:
		m.models.header.Focus()
		m.models.header.SelectSummary()
	case bodyComponent:
		m.models.body.Focus()
	}

	cmds := make([]tea.Cmd, 5)
	m.models.info, cmds[0] = m.models.info.Update(msg)
	m.models.header, cmds[1] = m.models.header.Update(msg)
	m.models.body, cmds[2] = m.models.body.Update(msg)
	m.models.footer, cmds[3] = m.models.footer.Update(msg)
	m.models.status, cmds[4] = m.models.status.Update(msg)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if m.err != nil {
		return fmt.Sprintf("unable to render view: %s", m.err)
	}

	return lipgloss.JoinVertical(lipgloss.Top,
		m.models.info.View(),
		m.models.header.View(),
		m.models.body.View(),
		m.models.footer.View(),
		m.models.status.View(),
	)
}

func (m Model) validate() bool {
	//nolint:gocritic
	switch {
	case m.result.Summary == "":
		return false
	}
	return true
}
