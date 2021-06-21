package handlers

import (
  "PasswordManager/inputs"
  "PasswordManager/services"
  "github.com/gin-gonic/gin"
  "net/http"
)

type userHandler struct {
  userService services.UserService
}

func NewUserHandler(userService services.UserService) *userHandler {
  return &userHandler{userService}
}

func (h *userHandler) SignUpUser(c *gin.Context) {
  var inputUserSignUp inputs.InputUserSignUp

  if err := c.ShouldBindJSON(&inputUserSignUp); err != nil {
    c.JSON(http.StatusUnprocessableEntity, gin.H{"errors" : err.Error()})
    return
  }

  userData, err := h.userService.SignUpUser(inputUserSignUp)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
    return
  }

  c.JSON(http.StatusOK, userData)
}