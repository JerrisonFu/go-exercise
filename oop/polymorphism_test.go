package oop

import (
	"math"
	"testing"
)

func TestRectangleArea(t *testing.T) {
	r := Rectangle{Width: 4, Height: 5}
	if r.Area() != 20 {
		t.Errorf("Rectangle Area = %f, want 20", r.Area())
	}
}

func TestRectanglePerimeter(t *testing.T) {
	r := Rectangle{Width: 4, Height: 5}
	if r.Perimeter() != 18 {
		t.Errorf("Rectangle Perimeter = %f, want 18", r.Perimeter())
	}
}

func TestCircleArea(t *testing.T) {
	c := Circle{Radius: 1}
	expected := math.Pi
	if math.Abs(c.Area()-expected) > 0.001 {
		t.Errorf("Circle Area = %f, want ~%f", c.Area(), expected)
	}
}

func TestCirclePerimeter(t *testing.T) {
	c := Circle{Radius: 1}
	expected := 2 * math.Pi
	if math.Abs(c.Perimeter()-expected) > 0.001 {
		t.Errorf("Circle Perimeter = %f, want ~%f", c.Perimeter(), expected)
	}
}

func TestTotalArea(t *testing.T) {
	shapes := []Shape{
		Rectangle{Width: 2, Height: 3},
		Circle{Radius: 1},
	}
	expected := 6.0 + math.Pi
	result := TotalArea(shapes)
	if math.Abs(result-expected) > 0.001 {
		t.Errorf("TotalArea = %f, want ~%f", result, expected)
	}
}

func TestTotalAreaEmpty(t *testing.T) {
	shapes := []Shape{}
	if TotalArea(shapes) != 0 {
		t.Error("TotalArea for empty shapes should be 0")
	}
}

func TestNewRectangle(t *testing.T) {
	r := NewRectangle(3, 4)
	if r.Width != 3 || r.Height != 4 {
		t.Errorf("NewRectangle(3,4) = (%f, %f), want (3, 4)", r.Width, r.Height)
	}
}

func TestNewCircle(t *testing.T) {
	c := NewCircle(5)
	if c.Radius != 5 {
		t.Errorf("NewCircle(5) = %f, want 5", c.Radius)
	}
}
