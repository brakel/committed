package theme

import (
	"sync"

	tea "github.com/charmbracelet/bubbletea"
	tint "github.com/lrstanley/bubbletint"
)

type Msg int

var (
	registry *tint.Registry
	once     sync.Once
)

var tints = []tint.Tint{
	tint.TintBuiltinDark,
	tint.TintDracula,
	tint.TintGithub,
	tint.TintMaterial,
	tint.TintMonokaiPro,
	tint.TintNord,
	tint.TintOneDark,
	tint.TintTokyoNight,
	tint.TintTomorrowNight,
}

func Tint() *tint.Registry {
	once.Do(func() {
		registry = tint.NewRegistry(tints[0], tints[1:]...)
	})

	return registry
}

//nolint:ireturn
func NextTint() tea.Msg {
	var msg Msg

	len := len(registry.TintIDs())
	ids := registry.TintIDs()

	if registry.ID() == ids[len-1] {
		registry.SetTintID(ids[0])
		return msg
	}

	registry.NextTint()

	return msg
}
