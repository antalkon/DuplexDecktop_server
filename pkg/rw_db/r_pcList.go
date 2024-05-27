package rw_db

import (
	"database/sql"
)

type PCLIST struct {
	ID       string `json:"id"`
	PcName   string `json:"pc_name"`
	PcStatus string `json:"pc_status"`
	PcIP     string `json:"pc_ip"`
}

func GetPcsFromDB(db *sql.DB) ([]PCLIST, error) {
	rows, err := db.Query("SELECT id, pc_name, pc_status, pc_ip FROM pc_base")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pcs []PCLIST
	for rows.Next() {
		var pc PCLIST
		err := rows.Scan(&pc.ID, &pc.PcName, &pc.PcStatus, &pc.PcIP)
		if err != nil {
			return nil, err
		}
		pcs = append(pcs, pc)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return pcs, nil
}
