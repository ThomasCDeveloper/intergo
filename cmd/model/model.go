package model

type mode int

const (
	boot mode = iota
	menu
	chat
	tooSmall
)

type Model struct {
	mode     mode
	lastMode mode

	appWidth int

	viewportWidth  int
	viewportHeight int
}

func DefaultModel() Model {
	return Model{
		appWidth: 80,
	}
}
