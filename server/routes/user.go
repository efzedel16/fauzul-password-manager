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
	authService    = auth.NewAuthService()
	userRepository = repositories.NewUserRepository(db)
	userService    = services.NewUserService(authService, userRepository)
	userHandler    = handlers.NewUserHandler(userService)
	authMiddle     = handlers.AuthMiddle(authService, userService)
)

func User(r *gin.Engine) {
	r.POST("/signup", userHandler.SignUp)
	r.POST("/signin", userHandler.SignIn)
	r.GET("/users", authMiddle, userHandler.GetAll)
	r.GET("/user/:user_id", authMiddle, userHandler.GetById)
}
