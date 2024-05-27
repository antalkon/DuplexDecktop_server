package handler

import (
	"github.com/antalkon/DuplexDecktop_srver/pkg/rw_db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) EditPc(c *gin.Context) {
	db, err := rw_db.ConnectDb()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	var pc rw_db.PC
	if err := c.ShouldBindJSON(&pc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pcId := pc.ID

	// Проверяем наличие ПК с заданным id
	exists, err := rw_db.CheckPCbyID(db, pcId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "PC not found"})
		return
	}

	// Обновляем информацию о ПК
	if err := rw_db.UpdatePC(db, pc); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "PC updated successfully"})
}
