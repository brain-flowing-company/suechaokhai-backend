package users

import (
	"github.com/brain-flowing-company/pprp-backend/internal/models"
	"gorm.io/gorm"
)

type Repository interface {
	GetAllUsers(*[]models.Users) error
	GetUserById(*models.Users, string) error
	GetUserByEmail(*models.Users, string) error
	CreateUser(*models.Users) error
	UpdateUser(*models.Users, string) error
	DeleteUser(string) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		db,
	}
}

func (repo *repositoryImpl) GetAllUsers(users *[]models.Users) error {
	return repo.db.Find(users).Error
}

func (repo *repositoryImpl) GetUserById(user *models.Users, userId string) error {
	return repo.db.First(&user, "user_id = ?", userId).Error
}

func (repo *repositoryImpl) GetUserByEmail(user *models.Users, email string) error {
	return repo.db.First(&user, "email = ?", email).Error
}

func (repo *repositoryImpl) CreateUser(user *models.Users) error {
	return repo.db.Create(&user).Error
}

func (repo *repositoryImpl) UpdateUser(user *models.Users, userId string) error {
	return repo.db.Where("user_id = ?", userId).Updates(&user).Error
}

func (repo *repositoryImpl) DeleteUser(userId string) error {
	return repo.db.Where("user_id = ?", userId).Delete(&models.Users{}).Error
}
