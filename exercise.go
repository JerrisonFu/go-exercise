package main

import (
	"exercise/basics"
	"exercise/concurrency"
	"exercise/database"
	"exercise/functions"
	"exercise/network"
	"exercise/oop"
	"exercise/stdlib"
	"exercise/web"
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

	println("\n=== Database ===")
	database.RunAll()

	println("\n=== Web (Gin) ===")
	web.RunAll()
}
