package handlers

import (
  "PasswordManager/auth"
  "PasswordManager/inputs"
  "PasswordManager/services"
  "github.com/gin-gonic/gin"
  "net/http"
)

type userHandler struct {
  authService auth.AuthService
  userService services.UserService
}

func NewUserHandler(authService auth.AuthService, userService services.UserService) *userHandler {
  return &userHandler{authService, userService}
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