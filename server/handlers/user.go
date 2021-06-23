package handlers

import (
	"FauzulPasswordManager/inputs"
	"FauzulPasswordManager/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) SignUp(c *gin.Context) {
	var input inputs.SignUp

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	data, err := h.userService.SignUp(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *userHandler) SignIn(c *gin.Context) {
	var input inputs.SignIn

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	data, err := h.userService.SignIn(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *userHandler) GetAll(c *gin.Context) {
	datas, err := h.userService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, datas)
}

func (h *userHandler) GetById(c *gin.Context) {
	id := c.Params.ByName("user_id")
	data, err := h.userService.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
