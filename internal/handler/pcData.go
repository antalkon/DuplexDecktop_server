package handler

import (
	"github.com/antalkon/DuplexDecktop_srver/pkg/rw_db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handler *Handler) PcData(c *gin.Context) {
	// Получение значения :id из маршрута
	id := c.Param("id")

	// Подключение к базе данных
	db, err := rw_db.ConnectDb()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	// Получение данных о ПК из базы данных
	pcInfo, err := rw_db.PcInfo(db, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if pcInfo == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "PC not found"})
		return
	}

	// Отправка данных о ПК в формате JSON
	c.JSON(http.StatusOK, pcInfo)
}
