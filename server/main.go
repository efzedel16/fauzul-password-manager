package main

import (
  "PasswordManager/routes"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  routes.UserRoute(r)

  if err := r.Run(); err != nil {
    return
  }
}