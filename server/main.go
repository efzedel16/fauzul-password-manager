package main

import (
  "FauzulPasswordManager/routes"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  routes.User(r)

  if err := r.Run(); err != nil {
    return
  }
}