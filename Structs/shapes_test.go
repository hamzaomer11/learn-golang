package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	expected := 40.0

	if got != expected {
		t.Errorf("got %.2f expected %.2f", got, expected)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape    Shape
		expected float64
	}{
		{shape: Rectangle{Width: 12.0, Height: 6.0}, expected: 72.0},
		{shape: Circle{Radius: 10}, expected: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, expected: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()

		if got != tt.expected {
			t.Errorf("got %g expected %g", got, tt.expected)
		}
	}

}
