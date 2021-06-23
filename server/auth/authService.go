package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
)

var key = os.Getenv("JWT_SECRET")

type authService struct {
}

func NewAuthService() *authService {
	return &authService{}
}

type AuthService interface {
	Generate(id int) (string, error)
	Validation(encoded string) (*jwt.Token, error)
}

func (s *authService) Generate(id int) (string, error) {
	payload := jwt.MapClaims{
		"user_id": id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, payload)
	signature, err := token.SignedString([]byte(key))
	if err != nil {
		return signature, err
	}

	return signature, nil
}

func (s *authService) Validation(encoded string) (*jwt.Token, error) {
	token, err := jwt.Parse(encoded, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}

		return key, nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
