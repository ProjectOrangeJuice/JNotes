package sql

import (
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

func GetNotes() []string {
	t := []string{"one", "two"}
	return t
}
