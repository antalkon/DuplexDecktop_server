package rw_db

import (
	"database/sql"
	"fmt"
)

func DelPCClass(db *sql.DB, id string) error {
	query := `
        UPDATE pc_base
        SET class = ''
        WHERE id = $1;
    `
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to update PC with id %s: %w", id, err)
	}
	return nil
}
