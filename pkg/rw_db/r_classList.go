package rw_db

import (
	"database/sql"
)

type CLASSLIST struct {
	ID          string         `json:"id"`
	ClassName   string         `json:"class_name"`
	ClassStatus string         `json:"class_status"`
	ClassShed   sql.NullString `json:"class_shedule"`
}

func GetClsFromDB(db *sql.DB) ([]CLASSLIST, error) {
	rows, err := db.Query("SELECT id, class_name, class_status, class_shedule FROM class_base")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pcs []CLASSLIST
	for rows.Next() {
		var pc CLASSLIST
		err := rows.Scan(&pc.ID, &pc.ClassName, &pc.ClassStatus, &pc.ClassShed)
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
