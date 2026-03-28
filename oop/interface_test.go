package oop

import (
	"slices"
	"testing"
)

func TestDogSpeak(t *testing.T) {
	d := Dog{Name: "Rex"}
	if d.Speak() != "Woof!" {
		t.Error("Dog should say Woof!")
	}
}

func TestCatSpeak(t *testing.T) {
	c := Cat{Name: "Mittens"}
	if c.Speak() != "Meow!" {
		t.Error("Cat should say Meow!")
	}
}

func TestDogMove(t *testing.T) {
	d := Dog{Name: "Rex"}
	if d.Move() != "runs" {
		t.Error("Dog should run")
	}
}

func TestCatMove(t *testing.T) {
	c := Cat{Name: "Mittens"}
	if c.Move() != "walks" {
		t.Error("Cat should walk")
	}
}

func TestMakeThemSpeak(t *testing.T) {
	animals := []Animal{Dog{Name: "Buddy"}, Cat{Name: "Whiskers"}}
	sounds := MakeThemSpeak(animals)
	expected := []string{"Woof!", "Meow!"}
	if !slices.Equal(sounds, expected) {
		t.Errorf("MakeThemSpeak() = %v, want %v", sounds, expected)
	}
}

func TestEmptyAnimals(t *testing.T) {
	animals := []Animal{}
	sounds := MakeThemSpeak(animals)
	if len(sounds) != 0 {
		t.Errorf("Expected empty slice, got %v", sounds)
	}
}
