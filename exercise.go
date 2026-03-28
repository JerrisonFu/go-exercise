package main

import (
	"exercise/basics"
	"exercise/concurrency"
	"exercise/functions"
	"exercise/oop"
)

func main() {
	println("=== Basics ===")
	basics.RunAll()

	println("\n=== Functions ===")
	functions.RunAll()

	println("\n=== OOP ===")
	oop.RunAll()

	println("\n=== Concurrency ===")
	concurrency.RunAll()
}
