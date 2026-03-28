package oop

type Animal interface {
	Speak() string
	Move() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

func (d Dog) Move() string {
	return "runs"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow!"
}

func (c Cat) Move() string {
	return "walks"
}

func MakeThemSpeak(animals []Animal) []string {
	results := make([]string, len(animals))
	for i, a := range animals {
		results[i] = a.Speak()
	}
	return results
}
