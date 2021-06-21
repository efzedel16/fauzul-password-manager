package repositories

import (
  "PasswordManager/entities"
  "gorm.io/gorm"
)

type userRepository struct {
  userDb *gorm.DB
}

func NewUserRepository(userDb *gorm.DB) *userRepository {
  return &userRepository{userDb}
}

type UserRepository interface {
  CreateUser(userData entities.User) (entities.User, error)
  UpdateUser(userId int, userDataUpdate map[string]interface{}) (entities.User, error)
  DeleteUser(userId int) (string, error)
  GetAllUsers() ([]entities.User, error)
  GetUser(userId int) (entities.User, error)
  FindUserById(userId int) (entities.User, error)
  FindUserByEmail(userEmail string) (entities.User, error)
}

func (r *userRepository) CreateUser(userData entities.User) (entities.User, error) {
  if err := r.userDb.Create(&userData).Error; err != nil {
    return userData, err
  }

  return userData, nil
}

func (r *userRepository) UpdateUser(userId int, userDataUpdate map[string]interface{}) (entities.User, error) {
  var userData entities.User
  if err := r.userDb.Where("id = ?", userId).Error; err != nil {
    return userData, err
  }

  if err := r.userDb.Model(&userData).Where("id = ?", userId).Updates(&userDataUpdate).Error; err != nil {
    return userData, err
  }

  return userData, nil
}

func (r *userRepository) DeleteUser(userId int) (string, error) {
  var userData entities.User
  if  err :=  r.userDb.Where("id = ?", userId).Delete(&userData).Error; err != nil {
    return "error", err
  }

  return "success", nil
}

func (r *userRepository) GetAllUsers() ([]entities.User, error) {
  var userDAta []entities.User
  if
}

func (r *userRepository) FindUserById(userId int) (entities.User, error) {
  var userData entities.User
  if  err := r.userDb.Where("id = ?", userId).Error; err != nil {
    return userData, err
  }

  return userData, nil
}

func (r *userRepository) FindUserByEmail(userEmail string) (entities.User, error) {
  var userData entities.User
  if err := r.userDb.Where("email = ?", userEmail).Error; err != nil {
    return userData, err
  }
  
  return userData, nil
}