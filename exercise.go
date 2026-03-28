package main

import (
	"exercise/basics"
	"exercise/concurrency"
	"exercise/functions"
	"exercise/network"
	"exercise/oop"
	"exercise/stdlib"
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

	println("\n=== Stdlib ===")
	stdlib.RunAll()

	println("\n=== Network ===")
	network.RunAll()
}
