package rw_db

import "database/sql"

type PcOffData struct {
	IP string `json:"ip"`
}

func PcOff(db *sql.DB, id string) (*PcOffData, error) {
	query := `
        SELECT pc_ip
        FROM pc_base
        WHERE id = $1
    `
	var pc PcOffData
	err := db.QueryRow(query, id).Scan(&pc.IP)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Если не найдено ни одной строки, возвращаем nil
		}
		return nil, err
	}

	return &pc, nil
}
