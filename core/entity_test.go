package core

import (
	"fmt"
	"testing"
)

type Square2 struct {
	*Square
	delta float64
}

var renderCalled bool = false

func (s *Square2) Render(d float64) error {
	renderCalled = true
	s.delta = d
	if s.width != 64 {
		return fmt.Errorf("ERROR %d", s.width)
	}

	return nil
}

func TestSquare(t *testing.T) {
	t.Run("width and height", func(t *testing.T) {
		s := &Square{100, 64}
		if s.width != 100 || s.height != 64 {
			t.Error("width and height should be 100 and 64")
		}
	})

	t.Run("as Entity", func(t *testing.T) {
		s := &Square{100, 64}
		var e Entity = s
		r1 := e.Render(0.1)
		r2 := e.Update(0.1)

		if r1 != nil || r2 != nil {
			t.Error("Render and Update should return nil")
		}

		s1, ok := e.(*Square)

		if !ok {
			t.Error("Wrong conversion")
		}

		s.width = 128

		if s1.width != 128 {
			t.Error("Square width should be 128")
		}
	})

	t.Run("overwrite Render method", func(t *testing.T) {
		s := &Square2{Square: &Square{64, 64}}
		err := s.Render(0.1)

		if !renderCalled {
			t.Error("Render wasn't called")
		}

		s.width = 24
		var e Entity = s
		err = e.Render(145)

		if err.Error() != "ERROR 24" {
			t.Error(err)
		}

		if s.delta != 145 {
			t.Errorf("Wrong delta %f", s.delta)
		}
	})
}
