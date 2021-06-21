package auth

import (
  "github.com/dgrijalva/jwt-go"
  "os"
)

var secretKey  = os.Getenv("JWT_SECRET")

type authService struct {
}

func NewAuthService() *authService {
  return &authService{}
}

type AuthService interface {
  GenerateToken(userId int) (string, error)
  TokenValidation(encodedToken string) (*jwt.Token, error)
}

func (s *authService) GenerateToken(userId int) (string, error) {

}

TokenValidation(encodedToken string) (*jwt.Token, error)