package handlers

import (
  "FauzulPasswordManager/inputs"
  "FauzulPasswordManager/services"
  "github.com/gin-gonic/gin"
  "net/http"
)

type handler struct {
  service services.UserService
}

func NewUserHandler(service services.UserService) *handler {
  return &handler{service}
}

func (h *handler) SignUp(c *gin.Context) {
  var input inputs.SignUp

  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusUnprocessableEntity, gin.H{"errors" : err.Error()})
    return
  }

  data, err := h.service.SignUp(input)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
    return
  }

  c.JSON(http.StatusOK, data)
}

func (h *handler) SignIn(c *gin.Context) {
  var input inputs.SignIn

  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusUnprocessableEntity, gin.H{"error" : err.Error()})
    return
  }

  data, err := h.service.SignIn(input)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
    return
  }

  c.JSON(http.StatusOK, data)
}

//func (h *handler) Update(id int, input inputs.UpdateUser) (entities.User, error) {
//
//}
//
//func (h *handler) Delete(id int) (string, error) {
//
//}
//
//func (h *handler) GetAll() ([]formatters.UserFormatter, error) {
//
//}