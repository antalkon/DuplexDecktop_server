package rw_db

import (
	"database/sql"
	"fmt"
)

type CLASS struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Status  string `json:"status"`
	Shedule string `json:"shedule"`
}

func CheckClassbyID(db *sql.DB, id string) (bool, error) {
	query := `
        SELECT EXISTS (
            SELECT 1
            FROM class_base
            WHERE id = $1
        )
    `
	var exists bool
	err := db.QueryRow(query, id).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check Class with id %s: %w", id, err)
	}
	return exists, nil
}
func UpdateClass(db *sql.DB, pc CLASS) error {
	query := `
        UPDATE class_base
        SET class_name = $1, class_status = $2, class_shedule = $3
        WHERE id = $4
    `
	_, err := db.Exec(query, pc.Name, pc.Status, pc.Shedule, pc.ID)
	if err != nil {
		return fmt.Errorf("failed to update PC with id %s: %w", pc.ID, err)
	}
	return nil
}
