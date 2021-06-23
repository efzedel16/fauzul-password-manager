package handlers

import (
	"FauzulPasswordManager/inputs"
	"FauzulPasswordManager/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type passHandler struct {
	passService services.PassService
}

func NewPassHandler(passService services.PassService) *passHandler {
	return &passHandler{passService}
}

func (h *passHandler) Create(c *gin.Context) {
	var input inputs.CreatePass

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	userId := c.MustGet("current_ser").(int)
	data, err := h.passService.Create(userId, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

//func (h *passHandler) Update(c *gin.Context) {
//  var input inputs.UpdatePass
//  id := c.Params.
//
//  if err := c.ShouldBindJSON(&input); err != nil {
//    c.JSON(http.StatusUnprocessableEntity, gin.H{"errors" : err.Error()})
//    return
//  }
//
//  data, err := h.passService.Update(id, input)
//  if err != nil {
//    c.JSON(http.StatusBadRequest, gin.H{"errors" : err.Error()})
//    return
//  }
//
//  c.JSON(http.StatusOK, data)
//}

//func (h *passHandler) Delete(c *gin.Context) {
//  id := c.Params.ByName("pass_id")
//
//  msg, err := h.passService.Delete(id)
//  if err != nil {
//    c.JSON(http.StatusBadRequest, gin.H{"errors" : err.Error()})
//    return
//  }
//
//  c.JSON(http.StatusOK, msg)
//}

//func (h *passHandler) GetAllByUserId(c *gin.Context) {
//  currentUser := c.MustGet("current_user").(entities.User)
//  //userId := strconv.Itoa(currentUser)
//  data, err := h.passService.GetAllByUserId(currentUser)
//  if err != nil {
//    c.JSON(http.StatusBadRequest, gin.H{"errors" : err.Error()})
//    return
//  }
//
//  if currentUser.Id == 0 {
//    c.JSON(http.StatusBadRequest, gin.H{"errors" : err.Error()})
//    return
//  }
//
//  c.JSON(http.StatusOK, data)
//}

//func (h *passHandler) GetById(c *gin.Context) {
//
//}
