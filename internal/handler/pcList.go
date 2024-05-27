package handler

import (
	"github.com/antalkon/DuplexDecktop_srver/pkg/rw_db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) ListPc(c *gin.Context) {
	db, err := rw_db.ConnectDb()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	pcs, err := rw_db.GetPcsFromDB(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pcs)
}
