package users

import (
	"github.com/brain-flowing-company/pprp-backend/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	GetAllUsers(*[]models.Users) error
	GetUserById(*models.Users, string) error
	CreateUser(*models.Users) error
	UpdateUser(*models.Users, string) error
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
	return repo.db.First(user, "user_id = ?", userId).Error
}

func (repo *repositoryImpl) CreateUser(user *models.Users) error {
	user.UserId = uuid.New()

	for repo.db.Find(&models.Users{}, "user_id = ?", user.UserId).RowsAffected != 0 {
		user.UserId = uuid.New()
	}

	return repo.db.Create(&user).Error
}

func (repo *repositoryImpl) UpdateUser(user *models.Users, userId string) error {
	return repo.db.Where("user_id = ?", userId).Updates(&user).Error
}
