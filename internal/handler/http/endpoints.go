package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/", h.addUser)
	router.GET("/info", nil)

	return router
}
