package basics

import "fmt"

func ControlFlow() {
	if 10 > 5 {
		fmt.Println("10 > 5")
	}

	score := 85
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		if i == 4 {
			break
		}
		fmt.Printf("for: %d\n", i)
	}

	counter := 0
	for counter < 3 {
		fmt.Printf("while: %d\n", counter)
		counter++
	}

	fruits := []string{"apple", "banana", "orange"}
	for index, fruit := range fruits {
		fmt.Printf("Index: %d, Fruit: %s\n", index, fruit)
	}

	switch day := "Monday"; day {
	case "Monday":
		fmt.Println("Start of work week")
	case "Friday":
		fmt.Println("End of work week")
	default:
		fmt.Println("Midweek")
	}
}
