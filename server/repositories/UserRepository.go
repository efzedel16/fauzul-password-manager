package repositories

import (
  //"PasswordManager/entities"
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
  UpdateUser(Id int, userDataUpdate map[string]interface{} (entities.User, error)
  DeleteUserUser(Id int) (string, error)
  FindById(Id int) (entities.User, error)
  FindByEmail(userEmail string) (entities.User, error)
}

func (r *userRepository) CreateUser(userData entities.User) (entities.User, error) {
  if err := r.userDb.Create(&userData).Error; err := nil {
    return user, err
  }

  return user, nil
}

func (r *userRepository) UpdateUser(Id int, userDataUpdate map[string]interface{} (entities.User, error) {
  var user entities.User
}

func (r *userRepository) DeleteUserUser(Id int) (string, error) {}
func (r *userRepository) FindById(Id int) (entities.User, error) {}
func (r *userRepository) FindByEmail(userEmail string) (entities.User, error) {}