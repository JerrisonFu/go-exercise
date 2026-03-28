package functions

func RunAll() {
	println("=== Basic Functions ===")
	println(Add(3, 5))
	println(Multiply(4, 6))

	println("\n=== Divide ===")
	result, err := Divide(10, 2)
	if err != nil {
		println("Error:", err.Error())
	} else {
		println("Result:", result)
	}
	_, err = Divide(5, 0)
	if err != nil {
		println("Expected error:", err.Error())
	}

	println("\n=== Variadic & Closure ===")
	println("Sum:", Sum(1, 2, 3, 4, 5))
	double := MakeMultiplier(2)
	println("Double 5:", double(5))

	println("\n=== Error Handling ===")
	err = ValidateAge(25)
	println("ValidateAge(25):", err)
	err = ValidateAge(-1)
	println("ValidateAge(-1):", err)

	println("\n=== Safe Divide (panic/recover) ===")
	println("SafeDivide(10, 2):", SafeDivide(10, 2))
	println("SafeDivide(5, 0):", SafeDivide(5, 0))
}
