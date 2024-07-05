package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("rectangles perimeter", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("circles perimeter", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Perimeter()
		want := 62.83185307179586

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

}

// table driven test
// A great fit when you wish to test various implementations of an interface or if the data being passed in to a function has lots of different requirements that need testing
func TestArea2(t *testing.T) {
	// anonymous struct: declaring a slice of structs by using []struct with 2 fields, the shape and the want
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		// {shape, want}
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}

}
