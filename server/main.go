package main

import (
	"FauzulPasswordManager/auth"
	"FauzulPasswordManager/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(auth.CORSMiddleware())

	routes.User(r)

	if err := r.Run(); err != nil {
		return
	}
}
