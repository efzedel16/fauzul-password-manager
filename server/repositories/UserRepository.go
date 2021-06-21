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
  DeleteUserUser(userId int) (string, error)
  FindById(userId int) (entities.User, error)
  FindByEmail(userEmail string) (entities.User, error)
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

//func (r *userRepository) DeleteUserUser(userId int) (string, error) {}
//func (r *userRepository) FindById(userId int) (entities.User, error) {}
//func (r *userRepository) FindByEmail(userEmail string) (entities.User, error) {}