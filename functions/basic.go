package functions

func Add(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisionByZero{}
	}
	return a / b, nil
}

type ErrDivisionByZero struct{}

func (e ErrDivisionByZero) Error() string {
	return "division by zero"
}
