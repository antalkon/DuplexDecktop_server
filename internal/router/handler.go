package router

import (
	"github.com/antalkon/DuplexDecktop_srver/internal/handler"
	"github.com/gin-gonic/gin"
)

type Handler struct{}

func (h *Handler) signUp(c *gin.Context) {
	// Реализуйте логику регистрации здесь
}

func (h *Handler) signIn(c *gin.Context) {
	// Реализуйте логику авторизации здесь
}

func (h *Handler) createList(c *gin.Context) {
	// Реализуйте логику создания списка здесь
}

func (h *Handler) getAllLists(c *gin.Context) {
	// Реализуйте логику получения всех списков здесь
}

func (h *Handler) getListById(c *gin.Context) {
	// Реализуйте логику получения списка по ID здесь
}

func (h *Handler) updateList(c *gin.Context) {
	// Реализуйте логику обновления списка здесь
}

func (h *Handler) deleteList(c *gin.Context) {
	// Реализуйте логику удаления списка здесь
}

func InitRouters(h *handler.Handler) *gin.Engine {
	router := gin.New()

	auth := router.Group("/pc")
	{
		auth.POST("/edit", h.AddNewPc)
		auth.PUT("/edit", h.EditPc)
		//auth.POST("/sign-in")
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.ListPc)
			//		lists.GET("/")
			lists.GET("/:id", h.PcData)
			lists.POST("/off/:id", h.PcOff)
			//		lists.PUT("/:id")
			//		lists.DELETE("/:id")
			//
		}
	}
	return router
}
