package users

import (
	"errors"

	"github.com/brain-flowing-company/pprp-backend/apperror"
	"github.com/brain-flowing-company/pprp-backend/internal/models"
	"github.com/brain-flowing-company/pprp-backend/utils"
	"gorm.io/gorm"
)

type Service interface {
	GetAllUsers(*[]models.Users) *apperror.AppError
	GetUserById(*models.Users, string) *apperror.AppError
	Register(*models.Users) *apperror.AppError
	UpdateUser(*models.Users, string) *apperror.AppError
	DeleteUser(string) *apperror.AppError
	GetUserByEmail(*models.Users, string) *apperror.AppError
}

type serviceImpl struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &serviceImpl{
		repo,
	}
}

func (service *serviceImpl) GetAllUsers(users *[]models.Users) *apperror.AppError {
	err := service.repo.GetAllUsers(users)

	if err != nil {
		return apperror.InternalServerError
	}

	return nil
}

func (s *serviceImpl) GetUserById(user *models.Users, userId string) *apperror.AppError {
	if !utils.IsValidUUID(userId) {
		return apperror.InvalidUserId
	}

	err := s.repo.GetUserById(user, userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return apperror.UserNotFound
	} else if err != nil {
		return apperror.InternalServerError
	}

	return nil
}

func (s *serviceImpl) Register(user *models.Users) *apperror.AppError {
	if !utils.IsValidEmail(user.Email) {
		return apperror.InvalidEmail
	}

	if s.repo.GetUserByEmail(&models.Users{}, user.Email) == nil {
		return apperror.EmailAlreadyExists
	}

	if user.Password != "" {
		if !utils.IsValidPassword(user.Password) {
			return apperror.InvalidPassword
		}

		hashedPassword, hashErr := utils.HashPassword(user.Password)
		if hashErr != nil {
			return apperror.PasswordCannotBeHashed
		}
		user.Password = string(hashedPassword)
	}

	err := s.repo.CreateUser(user)
	if err != nil {
		return apperror.InternalServerError
	}

	return nil
}

func (s *serviceImpl) UpdateUser(user *models.Users, userId string) *apperror.AppError {
	if !utils.IsValidUUID(userId) {
		return apperror.InvalidUserId
	}

	err := s.repo.UpdateUser(user, userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return apperror.UserNotFound
	} else if err != nil {
		return apperror.InternalServerError
	}

	return nil
}

func (s *serviceImpl) DeleteUser(userId string) *apperror.AppError {
	if !utils.IsValidUUID(userId) {
		return apperror.InvalidUserId
	}

	err := s.repo.DeleteUser(userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return apperror.UserNotFound
	} else if err != nil {
		return apperror.InternalServerError
	}

	return nil
}

func (s *serviceImpl) GetUserByEmail(user *models.Users, email string) *apperror.AppError {
	if !utils.IsValidEmail(email) {
		return apperror.InvalidEmail
	}

	err := s.repo.GetUserByEmail(user, email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return apperror.UserNotFound
	} else if err != nil {
		return apperror.InternalServerError
	}

	return nil
}
