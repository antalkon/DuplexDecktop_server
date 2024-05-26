package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) AddNewPc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "New PC added"})
}
