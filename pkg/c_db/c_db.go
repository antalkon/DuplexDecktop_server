package c_db

import (
	"fmt"
	"github.com/antalkon/DuplexDecktop_srver/internal/config"
	"log"
)

func DBLink() (string, error) {
	configData := config.MustDatabase()

	dbLink := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		configData.Host, configData.Port, configData.DBName, configData.User, configData.Password, configData.SSLMode)

	// Выводим дебаг-информацию
	log.Printf("DB link: %s\n", dbLink)

	return dbLink, nil
}
