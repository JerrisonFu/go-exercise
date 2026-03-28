package oop

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{Name: name, Age: age}
}

func (p Person) Greet() string {
	return "Hello, I'm " + p.Name
}

func (p *Person) SetAge(age int) {
	p.Age = age
}
