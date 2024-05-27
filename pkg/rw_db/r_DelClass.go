package rw_db

import (
	"database/sql"
	"fmt"
)

type DelClass struct {
	ID string `json:"id"`
}

func DeleteClass(db *sql.DB, delclass DelClass) error {
	query := `
        DELETE FROM class_base
        WHERE id = $1
    `
	_, err := db.Exec(query, delclass.ID)
	if err != nil {
		return fmt.Errorf("failed to delete class with id %s: %w", delclass.ID, err)
	}
	return nil
}
