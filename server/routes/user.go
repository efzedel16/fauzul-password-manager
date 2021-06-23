package routes

import (
	"FauzulPasswordManager/auth"
	"FauzulPasswordManager/config"
	"FauzulPasswordManager/handlers"
	"FauzulPasswordManager/repositories"
	"FauzulPasswordManager/services"
	"github.com/gin-gonic/gin"
)

var (
	db             = config.ConnectDb()
	userRepository = repositories.NewUserRepository(db)
	authService    = auth.NewAuthService()
	userService    = services.NewUserService(authService, userRepository)
	userHandler    = handlers.NewUserHandler(userService)
)

func User(r *gin.Engine) {
	r.POST("/signup", userHandler.SignUp)
	r.POST("/signin", userHandler.SignIn)
}
