package handler

import (
	"fmt"
	"net/http"
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

type People struct {
	Surname    string `json:"surname"`
	Name       string `json:"name"`
	Patronymic string `json:"patronymic,omitempty"`
	Address    string `json:"address"`
}

func getPeopleInfo(c *gin.Context) {
	var user domain.User

	user.PassportSerie = c.Query("passportSerie")
	user.PassportNumber = c.Query("passportNumber")

	// Здесь должна быть логика для поиска информации по пользователю
	// Например, запрос в базу данных для получения информации по переданным параметрам

	// Пример ответа
	people := People{
		Surname:    "Иванов",
		Name:       "Иван",
		Patronymic: "Иванович",
		Address:    "г. Москва, ул. Ленина, д. 5, кв. 1",
	}

	c.JSON(http.StatusOK, people)
}
