package shortcut

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type shortcutSet struct {
	align     int
	shortcuts []Shortcut
	modifiers []Modifier
	target    [][]string
	decorate  bool
	height    int
	width     int
	styles    Styles
}

func newShortcutSet(a int, ms []Modifier, ss []Shortcut, d bool) []string {
	scs := shortcutSet{
		align:     a,
		modifiers: ms,
		shortcuts: ss,
		decorate:  d,
		styles:    defaultStyles(),
	}
	scs.shortcutDimensions()
	scs.initShortcuts()
	scs.fillShortcuts()
	return scs.joinShortcuts()
}

func (s *shortcutSet) shortcutDimensions() {
	s.height = s.countModifiers()

	width := 0
	for _, v := range s.modifiers {
		if v.Align != s.align {
			continue
		}
		i := s.countShortcuts(v.Modifier)
		if i > width {
			width = i
		}
	}

	s.width = width
}

func (s *shortcutSet) initShortcuts() {
	col := make([][]string, s.height)

	for i := range col {
		row := make([]string, s.width*2)
		col[i] = row
	}

	s.target = col
}

func (s *shortcutSet) fillShortcuts() {
	i := 0
	for _, v := range s.modifiers {
		if v.Align != s.align {
			continue
		}

		j := 0
		for _, vv := range s.shortcuts {
			if v.Modifier != vv.Modifier {
				continue
			}

			switch s.align {
			case AlignLeft:
				s.target[i][j] = vv.Key
				s.target[i][j+1] = vv.Label
			case AlignRight:
				s.target[i][j+1] = vv.Key
				s.target[i][j] = vv.Label
			}

			j += 2
		}

		i++
	}
}

func (s *shortcutSet) joinShortcuts() []string {
	var ss []string
	var offset int

	if len(s.target) == 0 {
		return []string{}
	}

	for i := 0; i < len(s.target[0]); i++ {
		var col []string
		for j := 0; j < len(s.target); j++ {
			col = append(col, s.target[j][i])
		}
		len := sliceMaxLen(col)

		str := s.joinColumn(col, len, offset)

		if str == "" {
			continue
		}

		m := s.styles.shortcutColumnLeft.Render(str)
		if s.align == AlignRight {
			m = s.styles.shortcutColumnRight.Render(str)
		}

		ss = append(ss, m)

		offset++
	}
	return ss
}

func (s *shortcutSet) joinColumn(col []string, len int, offset int) string {
	var res []string

	remainder := 0
	lr := lipgloss.Right
	ll := lipgloss.Left

	if s.align == AlignRight {
		remainder = 1
		lr = lipgloss.Left
		ll = lipgloss.Right
	}

	for _, v := range col {
		switch offset%2 == remainder {
		case true:
			d := s.decorateKey(v, len, lr, s.decorate)
			res = append(res, d)
		case false:
			d := s.decorateLabel(v, len, ll, s.decorate)
			res = append(res, d)
		}
	}
	return strings.Join(res, "\n")
}

func (s *shortcutSet) countModifiers() int {
	i := 0
	for _, m := range s.modifiers {
		if m.Align != s.align {
			continue
		}
		i++
	}
	return i
}

func (s *shortcutSet) countShortcuts(modifier int) int {
	i := 0
	for _, s := range s.shortcuts {
		if s.Modifier != modifier {
			continue
		}
		i++
	}

	return i
}

func (s shortcutSet) decorateKey(key string, len int, align lipgloss.Position, bracket bool) string {
	var k string
	padding := 0

	if key != "" {
		k = s.styles.shortcutKey.Render(key)
		if bracket {
			padding = 2
			k = fmt.Sprintf("%v%v%v",
				s.styles.shortcutAngleBracket.Render("<"),
				s.styles.shortcutKey.Render(key),
				s.styles.shortcutAngleBracket.Render(">"),
			)
		}
	}

	return lipgloss.NewStyle().Width(len + padding).Align(align).Render(k)
}

func (s shortcutSet) decorateLabel(label string, len int, align lipgloss.Position, colour bool) string {
	var l string
	if label != "" {
		l = label
		if colour {
			l = s.styles.shortcutLabel.Render(label)
		}
	}

	return lipgloss.NewStyle().Width(len).Align(align).Render(l)
}

func sliceMaxLen(ss []string) int {
	var i int
	for _, v := range ss {
		if lipgloss.Width(v) > i {
			i = lipgloss.Width(v)
		}
	}

	return i
}
