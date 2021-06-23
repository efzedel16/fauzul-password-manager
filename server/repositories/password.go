package repositories

import (
	"FauzulPasswordManager/entities"
	"gorm.io/gorm"
)

type passRepository struct {
	db *gorm.DB
}

func NewPassRepository(db *gorm.DB) *passRepository {
	return &passRepository{db}
}

type PassRepository interface {
	Create(data entities.Password) (entities.Password, error)
	Update(id string, dataUpdate map[string]interface{}) (entities.Password, error)
	Delete(id string) (string, error)
	FindAll() ([]entities.Password, error)
	FindById(id string) (entities.Password, error)
	FindByUserId(userId string) ([]entities.Password, error)
}

func (r *passRepository) Create(data entities.Password) (entities.Password, error) {
	if err := r.db.Create(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (r *passRepository) Update(id string, dataUpdate map[string]interface{}) (entities.Password, error) {
	var data entities.Password

	if err := r.db.Where("id = ?", id).Error; err != nil {
		return data, err
	}

	if err := r.db.Model(&data).Where("id = ?", id).Updates(&dataUpdate).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (r *passRepository) Delete(id string) (string, error) {
	var data entities.User

	if err := r.db.Where("id = ?", id).Delete(&data).Error; err != nil {
		return "error", err
	}

	return "success", nil
}

func (r *passRepository) FindAll() ([]entities.Password, error) {
	var datas []entities.Password

	if err := r.db.Find(&datas).Error; err != nil {
		return datas, err
	}

	return datas, nil
}

func (r *passRepository) FindById(id string) (entities.Password, error) {
	var data entities.Password

	if err := r.db.Where("id = ?", id).Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (r *passRepository) FindByUserId(userId string) ([]entities.Password, error) {
	var data []entities.Password

	if err := r.db.Where("user_id", userId).Find(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}
