package services

import (
	"PasswordManager/entities"
	"PasswordManager/formatters"
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
	SignUpUser(InputUserSignUp inputs.InputUserSignUp) (formatters.UserFormatter, error)
	SignInUser(InputUserSignIn inputs.InputUserSignIn) (entities.User, error)
}

func (s *userService) SignUpUser(InputUserSignUp inputs.InputUserSignUp) (formatters.UserFormatter, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(InputUserSignUp.Password), bcrypt.DefaultCost)
	userData := entities.User{
		FullName:     InputUserSignUp.FullName,
		Address:      InputUserSignUp.Address,
		Email:        InputUserSignUp.Email,
		PasswordHash: string(passHash),
	}

	if err != nil {
		return formatters.UserFormatter{}, err
	}

	newUser, err := s.userRepository.CreateUser(userData)
	if err != nil {
		return formatters.UserFormatter{}, err
	}

	return formatters.UserFormat(newUser), nil
}

func (s *userService) SignInUser(InputUserSignIn inputs.InputUserSignIn) (entities.User, error) {
	userData, err := s.userRepository.FindUserByEmail(InputUserSignIn.Email)
	if err != nil {
		return userData, err
	}

	if userData.Id == 0 {
		return entities.User{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userData.PasswordHash), []byte(InputUserSignIn.Password)); err != nil {
		return userData, err
	}

	return userData, nil
}