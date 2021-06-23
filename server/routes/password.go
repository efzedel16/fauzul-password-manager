package routes

import (
	"FauzulPasswordManager/handlers"
	"FauzulPasswordManager/repositories"
	"FauzulPasswordManager/services"
	"github.com/gin-gonic/gin"
)

var (
	passRepository = repositories.NewPassRepository(db)
	passService    = services.NewPassService(passRepository)
	passHandler    = handlers.NewPassHandler(passService)
)

func Pass(r *gin.Engine) {
	r.POST("/password", authMiddle, passHandler.Create)
	r.PUT("/password/:pass_id", authMiddle, passHandler.Update)
	r.DELETE("/password/:pass_id", authMiddle, passHandler.Delete)
	r.GET("/passwords", authMiddle, passHandler.GetAllByUserId)
	r.GET("/password/:pass_id", authMiddle, passHandler.GetById)
}