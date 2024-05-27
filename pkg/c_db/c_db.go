package c_db

import (
	"fmt"
	"github.com/antalkon/DuplexDecktop_srver/internal/config"
)

func DBLink() (string, error) {
	configData := config.MustDatabase()

	dbLink := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		configData.Host, configData.Port, configData.DBName, configData.User, configData.Password, configData.SSLMode)

	return dbLink, nil
}
