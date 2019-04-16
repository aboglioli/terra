package entity

type Renderer interface {
	Render(delta float64) error
}

type Updater interface {
	Update(delta float64) error
}

type Entity interface {
	Renderer
	Updater
}

type Square struct {
	width  int
	height int
}

func (s *Square) Render(d float64) error {
	return nil
}

func (s *Square) Update(d float64) error {
	return nil
}
