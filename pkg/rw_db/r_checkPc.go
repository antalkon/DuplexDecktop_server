package rw_db

import (
	"database/sql"
	"fmt"
)

func CheckPCbyID(db *sql.DB, id string) (bool, error) {
	query := `
        SELECT EXISTS (
            SELECT 1
            FROM pc_base
            WHERE id = $1
        )
    `
	var exists bool
	err := db.QueryRow(query, id).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check PC with id %s: %w", id, err)
	}
	return exists, nil
}
