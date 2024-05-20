package structs

import "testing"

func TestPerimeter(t *testing.T) {
	r := Rectangle{12.0, 6.0}
	got := Perimeter(r)
	want := 36.0

	if got != want {
		t.Errorf("got %.2f, but want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12.0, 6.0}, hasArea: 72.0},
		{name: "Cirlce", shape: Circle{10.0}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %#g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}

	// checkArea := func(t testing.TB, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("got %g want %g", got, want)
	// 	}
	// }

	// t.Run("rectangles", func(t *testing.T) {
	// 	r := Rectangle{12, 6}
	// 	checkArea(t, r, 72.0)
	// })

	// t.Run("circles", func(t *testing.T) {
	// 	c := Circle{10.0}
	// 	checkArea(t, c, 314.1592653589793)
	// })
}
