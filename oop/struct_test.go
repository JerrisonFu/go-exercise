package oop

import (
	"testing"
)

func TestNewPerson(t *testing.T) {
	p := NewPerson("Bob", 25)
	if p.Name != "Bob" {
		t.Errorf("Expected Name 'Bob', got '%s'", p.Name)
	}
	if p.Age != 25 {
		t.Errorf("Expected Age 25, got %d", p.Age)
	}
}

func TestPersonGreet(t *testing.T) {
	p := Person{Name: "Charlie"}
	expected := "Hello, I'm Charlie"
	if p.Greet() != expected {
		t.Errorf("Greet() = '%s', want '%s'", p.Greet(), expected)
	}
}

func TestPersonSetAge(t *testing.T) {
	p := &Person{Name: "Dave", Age: 20}
	p.SetAge(25)
	if p.Age != 25 {
		t.Errorf("After SetAge(25), Age = %d, want 25", p.Age)
	}
}

func TestPersonMutation(t *testing.T) {
	p := Person{Name: "Eve", Age: 18}
	p.SetAge(19)
	if p.Age != 19 {
		t.Error("SetAge should mutate the value")
	}
}
