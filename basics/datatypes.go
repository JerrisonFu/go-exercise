package basics

import "fmt"

func DataTypes() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Array: %v\n", arr)

	slice := []string{"apple", "banana", "cherry"}
	fmt.Printf("Slice: %v\n", slice)

	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	fmt.Printf("Map: %v\n", m)

	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Tom", Age: 20}
	fmt.Printf("Struct: %+v\n", p)
}
