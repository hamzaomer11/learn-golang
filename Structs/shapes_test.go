package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	expected := 40.0

	if got != expected {
		t.Errorf("got %.2f expected %.2f", got, expected)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()
		got := shape.Area()

		if got != expected {
			t.Errorf("got %g expected %g", got, expected)
		}
	}

	t.Run("area of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("area of a circle", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}
