package handlers

import (
	"FauzulPasswordManager/auth"
	"FauzulPasswordManager/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AuthMiddle(authService auth.AuthService, userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")

		token, err := authService.Validation(header)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorize user"})
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorize user"})
			return
		}

		id := strconv.Itoa(int(claim["user_id"].(float64)))
		current, err := userService.GetById(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorize user"})
			return
		}

		c.Set("current_user", current)
	}

}
