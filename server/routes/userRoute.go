package routes

import (
  "PasswordManager/auth"
  "PasswordManager/config"
  "PasswordManager/hendlers"
  "PasswordManager/repositories"
  "PasswordManager/services"
  "github.com/gin-gonic/gin"
)

var (
  db             = config.ConnectDb()
  userRepository = repositories.NewUserRepository(db)
  userService    = services.NewUserService(userRepository)
  authService    = auth.NewAuthService()
  userHandler    = handlers.NewUserHandler(authService, userService)
  )

func UserRoute(r *gin.Engine)  {
  user := r.Group("user")
  user.POST("/signup", userHandler.SignUpUser)
}