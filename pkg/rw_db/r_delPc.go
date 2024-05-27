package rw_db

import (
	"database/sql"
	"fmt"
)

type DelPc struct {
	ID string `json:"id"`
}

func DeletePc(db *sql.DB, delpc DelPc) error {
	query := `
        DELETE FROM pc_base
        WHERE id = $1
    `
	_, err := db.Exec(query, delpc.ID)
	if err != nil {
		return fmt.Errorf("failed to delete class with id %s: %w", delpc.ID, err)
	}
	return nil
}
