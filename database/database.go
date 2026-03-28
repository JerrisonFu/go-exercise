package database

import (
	"fmt"
	"os"
)

func RunAll() {
	db, err := OpenDB("test.db")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer os.Remove("test.db")
	defer db.Close()

	fmt.Println("=== Create Table ===")
	if err := db.CreateTable(); err != nil {
		fmt.Println("Error creating table:", err)
		return
	}
	fmt.Println("Table created successfully")

	fmt.Println("\n=== Insert Users ===")
	id1, _ := db.Insert("Alice", 30)
	fmt.Printf("Inserted Alice, ID: %d\n", id1)
	id2, _ := db.Insert("Bob", 25)
	fmt.Printf("Inserted Bob, ID: %d\n", id2)

	fmt.Println("\n=== Query All Users ===")
	users, _ := db.QueryAll()
	for _, u := range users {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", u.ID, u.Name, u.Age)
	}

	fmt.Println("\n=== Query Single User ===")
	user, _ := db.QueryByID(id1)
	fmt.Printf("QueryByID(%d): %s, %d years old\n", id1, user.Name, user.Age)

	fmt.Println("\n=== Update User ===")
	db.Update(id1, "Alice Updated", 31)
	user, _ = db.QueryByID(id1)
	fmt.Printf("After update: %s, %d years old\n", user.Name, user.Age)

	fmt.Println("\n=== Delete User ===")
	db.Delete(id2)
	users, _ = db.QueryAll()
	fmt.Printf("After delete: %d user(s) remaining\n", len(users))

	fmt.Println("\n=== CRUD Complete ===")
}
