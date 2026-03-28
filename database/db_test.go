package database

import (
	"testing"
)

func TestCreateTable(t *testing.T) {
	db, err := OpenMemoryDB()
	if err != nil {
		t.Fatalf("OpenMemoryDB() error = %v", err)
	}
	defer db.Close()

	err = db.CreateTable()
	if err != nil {
		t.Errorf("CreateTable() error = %v", err)
	}
}

func TestInsert(t *testing.T) {
	db, err := OpenMemoryDB()
	if err != nil {
		t.Fatalf("OpenMemoryDB() error = %v", err)
	}
	defer db.Close()

	db.CreateTable()

	id, err := db.Insert("Alice", 30)
	if err != nil {
		t.Errorf("Insert() error = %v", err)
	}
	if id <= 0 {
		t.Errorf("Insert() returned invalid id = %d", id)
	}
}

func TestQueryByID(t *testing.T) {
	db, err := OpenMemoryDB()
	if err != nil {
		t.Fatalf("OpenMemoryDB() error = %v", err)
	}
	defer db.Close()

	db.CreateTable()
	id, _ := db.Insert("Bob", 25)

	user, err := db.QueryByID(id)
	if err != nil {
		t.Errorf("QueryByID() error = %v", err)
	}
	if user.Name != "Bob" {
		t.Errorf("user.Name = %s, want 'Bob'", user.Name)
	}
	if user.Age != 25 {
		t.Errorf("user.Age = %d, want 25", user.Age)
	}
}

func TestQueryAll(t *testing.T) {
	db, err := OpenMemoryDB()
	if err != nil {
		t.Fatalf("OpenMemoryDB() error = %v", err)
	}
	defer db.Close()

	db.CreateTable()
	db.Insert("Charlie", 30)
	db.Insert("David", 35)

	users, err := db.QueryAll()
	if err != nil {
		t.Errorf("QueryAll() error = %v", err)
	}
	if len(users) != 2 {
		t.Errorf("len(users) = %d, want 2", len(users))
	}
}

func TestUpdate(t *testing.T) {
	db, err := OpenMemoryDB()
	if err != nil {
		t.Fatalf("OpenMemoryDB() error = %v", err)
	}
	defer db.Close()

	db.CreateTable()
	id, _ := db.Insert("Eve", 20)

	err = db.Update(id, "Eve Updated", 21)
	if err != nil {
		t.Errorf("Update() error = %v", err)
	}

	user, _ := db.QueryByID(id)
	if user.Name != "Eve Updated" {
		t.Errorf("user.Name = %s, want 'Eve Updated'", user.Name)
	}
	if user.Age != 21 {
		t.Errorf("user.Age = %d, want 21", user.Age)
	}
}

func TestDelete(t *testing.T) {
	db, err := OpenMemoryDB()
	if err != nil {
		t.Fatalf("OpenMemoryDB() error = %v", err)
	}
	defer db.Close()

	db.CreateTable()
	id, _ := db.Insert("Frank", 40)

	err = db.Delete(id)
	if err != nil {
		t.Errorf("Delete() error = %v", err)
	}

	_, err = db.QueryByID(id)
	if err == nil {
		t.Error("QueryByID() should return error after delete")
	}
}

func TestCRUD(t *testing.T) {
	db, err := OpenMemoryDB()
	if err != nil {
		t.Fatalf("OpenMemoryDB() error = %v", err)
	}
	defer db.Close()

	db.CreateTable()

	id, _ := db.Insert("Grace", 28)
	user, _ := db.QueryByID(id)
	if user.Name != "Grace" || user.Age != 28 {
		t.Error("Insert/Query failed")
	}

	db.Update(id, "Grace H.", 29)
	user, _ = db.QueryByID(id)
	if user.Age != 29 {
		t.Error("Update failed")
	}

	all, _ := db.QueryAll()
	if len(all) != 1 {
		t.Error("QueryAll failed")
	}

	db.Delete(id)
	all, _ = db.QueryAll()
	if len(all) != 0 {
		t.Error("Delete failed")
	}
}
