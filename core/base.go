package core

type Entity interface {
	Update(delta float64) error
	Render(delta float64) error
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
