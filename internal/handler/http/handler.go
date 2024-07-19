package handler

import (
	"fmt"
	"test/internal/domain"
	"test/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{service: services}
}

func (h *Handler) addUser(c *gin.Context) {
	var user domain.User

	if err := c.BindJSON(&user); err != nil {
		fmt.Errorf("error BindJSON:%w", err)
	}
}
