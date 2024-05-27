package handler

import (
	"github.com/google/uuid"
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

	// Декодируем JSON данные из запроса
	var pc rw_db.PC
	if err := c.ShouldBindJSON(&pc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pcId := generateUUID()
	// Добавляем новый ПК в базу данных
	if err := rw_db.WriteNewPc(db, pc, pcId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем успешный ответ
	c.JSON(http.StatusOK, gin.H{"message": "New PC added"})
}

// ПЕРЕНОС
func generateUUID() string {
	uuidObj := uuid.New()
	return uuidObj.String()[:6] // Берем первые 6 символов UUID
}
