package services

import (
  "PasswordManager/entities"
  "PasswordManager/inputs"
  "PasswordManager/repositories"
  "golang.org/x/crypto/bcrypt"
)

type userService struct {
  userRepository repositories.UserRepository
}

func NewUserService(userRepository repository.userRepository) *userService {
  return &userService{userRepository}
}

type UserService interface {
  SignUpUser(inputSignUp inputs.InputSignUp) (entities.User, error)
}

func (s *userService) SignUpUser(inputSignUp inputs.InputUserSignUp) (entities.User, error) {
  passHash, err :=  bcrypt.GenerateFromPassword([])
}