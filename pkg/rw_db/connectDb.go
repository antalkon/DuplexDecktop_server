package rw_db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/antalkon/DuplexDecktop_srver/pkg/c_db"
	_ "github.com/lib/pq"
)

func ConnectDb() (*sql.DB, error) {
	psqlInfo, err := c_db.DBLink()
	if err != nil {
		return nil, fmt.Errorf("failed to get database connection string: %w", err)
	}
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Проверяем соединение
	err = db.Ping()
	if err != nil {
		db.Close() // Закрываем соединение в случае ошибки
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to the database!")
	return db, nil
}
