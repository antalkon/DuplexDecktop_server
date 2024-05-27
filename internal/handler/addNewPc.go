package handler

import (
	"github.com/antalkon/DuplexDecktop_srver/pkg/g_uuid"
	"net/http"

	"github.com/antalkon/DuplexDecktop_srver/pkg/rw_db"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AddNewPc(c *gin.Context) {
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

	pcId := g_uuid.GenerateUUID()

	if err := rw_db.WriteNewPc(db, pc, pcId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "New PC added"})
}
