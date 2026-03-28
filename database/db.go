package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type User struct {
	ID   int64
	Name string
	Age  int
}

type DB struct {
	*sql.DB
}

func OpenDB(path string) (*DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func (db *DB) CreateTable() error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			age INTEGER NOT NULL
		)
	`)
	return err
}

func (db *DB) Insert(name string, age int) (int64, error) {
	result, err := db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", name, age)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (db *DB) QueryByID(id int64) (*User, error) {
	user := &User{}
	err := db.QueryRow("SELECT id, name, age FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (db *DB) QueryAll() ([]*User, error) {
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		user := &User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, rows.Err()
}

func (db *DB) Update(id int64, name string, age int) error {
	_, err := db.Exec("UPDATE users SET name = ?, age = ? WHERE id = ?", name, age, id)
	return err
}

func (db *DB) Delete(id int64) error {
	_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}

func (db *DB) Close() error {
	return db.DB.Close()
}

func OpenMemoryDB() (*DB, error) {
	return OpenDB(":memory:")
}
