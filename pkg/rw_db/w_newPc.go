package rw_db

import (
	"database/sql"
	"fmt"
)

type PC struct {
	ID          string `json:"id"`
	PcName      string `json:"pc_name"`
	PcIP        string `json:"pc_ip"`
	PcStatus    string `json:"pc_status"`
	PcConfig    string `json:"pc_config"`
	PcLocalname string `json:"pc_localname"`
	Coment      string `json:"coment"`
}

func WriteNewPc(db *sql.DB, pc PC, id string) error {
	// Генерируем UUID

	query := `
        INSERT INTO pc_base (id, pc_name, pc_ip, pc_status, pc_config, pc_localname, coment)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
    `
	_, err := db.Exec(query, id, pc.PcName, pc.PcIP, pc.PcStatus, pc.PcConfig, pc.PcLocalname, pc.Coment)
	if err != nil {
		return fmt.Errorf("failed to insert new PC: %w", err)
	}
	return nil
}
