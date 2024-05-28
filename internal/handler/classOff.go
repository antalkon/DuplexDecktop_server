package handler

import (
	"github.com/antalkon/DuplexDecktop_srver/pkg/rw_db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handler *Handler) ClassOff(c *gin.Context) {
	id := c.Param("id")

	db, err := rw_db.ConnectDb()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pcData, err := rw_db.ClassOff(db, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Отправляем байтовый срез JSON в ответе
	c.Data(http.StatusOK, "application/json", pcData)
}
