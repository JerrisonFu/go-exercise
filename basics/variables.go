package basics

import "fmt"

func Variables() {
	var name string = "Go"
	var age int = 14
	var isCool bool = true
	var version float64 = 1.21

	fmt.Printf("Name: %s, Age: %d, IsCool: %t, Version: %.2f\n", name, age, isCool, version)

	x := 10
	y := 20
	fmt.Printf("x + y = %d\n", x+y)

	const PI = 3.14159
	fmt.Printf("PI = %f\n", PI)
}
