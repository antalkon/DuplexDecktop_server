package g_uuid

import "github.com/google/uuid"

// GenerateUUID генерирует UUID и возвращает его первые 6 символов
func GenerateUUID() string {
	uuidObj := uuid.New()
	return uuidObj.String()[:6] // Берем первые 6 символов UUID
}
