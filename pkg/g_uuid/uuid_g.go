package g_uuid

import "github.com/google/uuid"

func generateUUID() string {
	uuidObj := uuid.New()
	return uuidObj.String()[:6] // Берем первые 6 символов UUID
}
