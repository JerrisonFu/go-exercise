package functions

import (
	"errors"
	"fmt"
)

func ValidateAge(age int) error {
	if age < 0 {
		return errors.New("age cannot be negative")
	}
	if age > 150 {
		return errors.New("age is too large")
	}
	return nil
}

func ValidateAgeWithCustom(age int) error {
	if age < 0 {
		return ErrInvalidAge{Msg: "age cannot be negative"}
	}
	if age > 150 {
		return ErrInvalidAge{Msg: "age is too large"}
	}
	return nil
}

type ErrInvalidAge struct {
	Msg string
}

func (e ErrInvalidAge) Error() string {
	return e.Msg
}

func SafeDivide(a, b int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			result = 0
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()
	return a / b
}
