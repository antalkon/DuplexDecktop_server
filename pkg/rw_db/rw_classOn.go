package rw_db

import (
	"database/sql"
	"encoding/json"
)

func ClassOn(db *sql.DB, id string) ([]byte, error) {
	query := `
		SELECT pc_ip
		FROM pc_base
		WHERE class = $1
	`
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []Result
	for rows.Next() {
		var pcIP string
		if err := rows.Scan(&pcIP); err != nil {
			return nil, err
		}
		results = append(results, Result{PcIP: pcIP})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(results)
	if err != nil {
		return nil, err
	}

	query2 := `
        UPDATE class_base
        SET class_status = 'true'
        WHERE id = $1
        `
	_, err = db.Exec(query2, id)
	return jsonData, nil
}
