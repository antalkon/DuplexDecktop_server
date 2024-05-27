package rw_db

import (
	"database/sql"
	"fmt"
)

func UpdatePC(db *sql.DB, pc PC) error {
	query := `
        UPDATE pc_base
        SET pc_name = $1, pc_ip = $2, pc_status = $3, pc_config = $4, pc_localname = $5, coment = $6
        WHERE id = $7
    `
	_, err := db.Exec(query, pc.PcName, pc.PcIP, pc.PcStatus, pc.PcConfig, pc.PcLocalname, pc.Coment, pc.ID)
	if err != nil {
		return fmt.Errorf("failed to update PC with id %s: %w", pc.ID, err)
	}
	return nil
}
