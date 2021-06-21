package auth

import (
  "errors"
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
  userPayload := jwt.MapClaims{
    "user_id" : userId,
  }

  userToken := jwt.NewWithClaims(jwt.SigningMethodHS512, userPayload)
  signature, err := userToken.SignedString([]byte(secretKey))
  if err != nil {
    return signature, err
  }

  return signature, nil
}

func (s *authService) TokenValidation(encodedToken string) (*jwt.Token, error) {
  userToken, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, errors.New("wrong token")
    }

    return secretKey, nil
  })

  if err != nil {
    return userToken, err
  }

  return userToken, nil
}