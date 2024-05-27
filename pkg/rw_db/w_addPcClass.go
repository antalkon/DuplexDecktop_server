package rw_db

import (
	"database/sql"
	"fmt"
)

type AddPcClass struct {
	ID      string `json:"id"`
	ClassID string `json:"class"`
}

func AddPCClass(db *sql.DB, class AddPcClass, id string) error {
	query := `
		UPDATE pc_base
        SET class = $1
        WHERE id = $2
    
`
	_, err := db.Exec(query, class.ClassID, id)
	if err != nil {
		return fmt.Errorf("failed to update PC with id %s: %w", id, err)
	}
	return nil
}
