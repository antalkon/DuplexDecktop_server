package rw_db

import (
	"database/sql"
	"fmt"
)

type Class struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewClass(db *sql.DB, class Class, id string) error {
	query := `
        INSERT INTO class_base (id, class_name)
        VALUES ($1, $2)
    `
	_, err := db.Exec(query, id, class.Name)
	if err != nil {
		return fmt.Errorf("failed to insert new PC: %w", err)
	}
	return nil
}
