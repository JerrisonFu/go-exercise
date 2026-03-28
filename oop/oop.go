package oop

type Address struct {
	Street  string
	City    string
	ZipCode string
}

type Employee struct {
	Name    string
	Salary  float64
	Address Address
}

func (e Employee) GetInfo() string {
	return e.Name + " - " + e.Address.City
}

func RunAll() {
	println("=== Struct ===")
	p := NewPerson("Alice", 30)
	println(p.Greet())
	p.SetAge(31)
	println("Age updated:", p.Age)

	println("\n=== Interface ===")
	animals := []Animal{Dog{Name: "Buddy"}, Cat{Name: "Whiskers"}}
	sounds := MakeThemSpeak(animals)
	for i, s := range sounds {
		println(animals[i].Move(), "->", s)
	}

	println("\n=== Polymorphism ===")
	shapes := []Shape{
		Rectangle{Width: 4, Height: 5},
		Circle{Radius: 3},
	}
	for _, s := range shapes {
		println("Area:", s.Area(), "Perimeter:", s.Perimeter())
	}
	println("Total Area:", TotalArea(shapes))
}
