package routes

import (
	"PasswordManager/auth"
	"PasswordManager/config"
	handlers "PasswordManager/hendlers"
	"PasswordManager/repositories"
	"PasswordManager/services"
	"github.com/gin-gonic/gin"
)

var (
	db             = config.ConnectDb()
	userRepository = repositories.NewUserRepository(db)
	authService    = auth.NewAuthService()
	userService    = services.NewUserService(userRepository)
	userHandler    = handlers.NewUserHandler(authService, userService)
)

func UserRoute(r *gin.Engine) {
	user := r.Group("user")
	user.POST("/signup", userHandler.SignUpUser)
}
