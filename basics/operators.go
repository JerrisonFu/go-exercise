package basics

import "fmt"

func OperatorsAndPointers() {
	a, b := 10, 3
	fmt.Printf("a + b = %d\na - b = %d\na * b = %d\na / b = %d\na %% b = %d\n",
		a+b, a-b, a*b, a/b, a%b)

	x, y := 5, 10
	fmt.Printf("x == y: %t\nx != y: %t\nx < y: %t\nx > y: %t\n",
		x == y, x != y, x < y, x > y)

	num := 5
	ptr := &num
	fmt.Printf("num value: %d, address: %p\n", num, ptr)
	fmt.Printf("ptr value (dereference): %d\n", *ptr)

	*ptr = 20
	fmt.Printf("num after change via ptr: %d\n", num)
}
