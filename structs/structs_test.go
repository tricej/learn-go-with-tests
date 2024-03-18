package structs

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		r := Rectangle{10.0, 10.0}
		got := r.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		c := Circle{5.0}
		got := c.Perimeter()
		want := 31.41592653589793

		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})

}

func TestArea(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {
		r := Rectangle{12, 6}
		got := r.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		c := Circle{10}
		got := c.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

}
