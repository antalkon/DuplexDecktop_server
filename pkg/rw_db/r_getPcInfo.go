package rw_db

import "database/sql"

type PCINFO struct {
	ID          string `json:"id"`
	PcName      string `json:"pc_name"`
	PcConfig    string `json:"pc_config"`
	PcStatus    string `json:"pc_status"`
	PcIP        string `json:"pc_ip"`
	PcLocalname string `json:"pc_localname"`
	Coment      string `json:"coment"`
}

func PcInfo(db *sql.DB, id string) (*PCINFO, error) {
	query := `
        SELECT id, pc_name, pc_config, pc_status, pc_ip, pc_localname, coment
        FROM pc_base
        WHERE id = $1
    `
	var pc PCINFO
	err := db.QueryRow(query, id).Scan(&pc.ID, &pc.PcName, &pc.PcConfig, &pc.PcStatus, &pc.PcIP, &pc.PcLocalname, &pc.Coment)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Если не найдено ни одной строки, возвращаем nil
		}
		return nil, err
	}

	return &pc, nil
}
