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

func NewUserService(userRepository repositories.UserRepository) *userService {
  return &userService{userRepository}
}

type UserService interface {
  SignUpUser(InputUserSignUp inputs.InputUserSignUp) (entities.User, error)
}

func (s *userService) SignUpUser(InputUserSignUp inputs.InputUserSignUp) (entities.User, error) {
  passHash, err :=  bcrypt.GenerateFromPassword([]byte(InputUserSignUp.Password), bcrypt.DefaultCost)
  userData := entities.User{
    FullName: InputUserSignUp.FullName,
    Address: InputUserSignUp.Address,
    Email: InputUserSignUp.Email,
    PasswordHash: string(passHash),
  }

  if err != nil {
    return userData, err
  }

  newUser, err := s.userRepository.CreateUser(userData)
  if err != nil {
    return newUser, err
  }

  return newUser, nil
}