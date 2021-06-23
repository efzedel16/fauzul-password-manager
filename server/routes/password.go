package routes

import (
	"FauzulPasswordManager/handlers"
	"FauzulPasswordManager/repositories"
	"FauzulPasswordManager/services"
	"github.com/gin-gonic/gin"
)

var (
	passRepository = repositories.NewPassRepository(db)
	passService    = services.NewPassService(authService, passRepository)
	passHandler    = handlers.NewPassHandler(passService)
)

func Pass(r *gin.Engine) {
	r.POST("/create", passHandler.Create)
	//r.POST("/signin", passHandler.SignIn)
}
