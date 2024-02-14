package users

import (
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/brain-flowing-company/pprp-backend/apperror"
	"github.com/brain-flowing-company/pprp-backend/config"
	"github.com/brain-flowing-company/pprp-backend/internal/models"
	"github.com/brain-flowing-company/pprp-backend/storage"
	"github.com/brain-flowing-company/pprp-backend/utils"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service interface {
	GetAllUsers(*[]models.Users) *apperror.AppError
	GetUserById(*models.Users, string) *apperror.AppError
	Register(*models.Users, models.Session) *apperror.AppError
	UpdateUser(*models.Users, string) *apperror.AppError
	DeleteUser(string) *apperror.AppError
	GetUserByEmail(*models.Users, string) *apperror.AppError
	UploadProfileImage(uuid.UUID, string, io.Reader) (string, *apperror.AppError)
}

type serviceImpl struct {
	repo    Repository
	logger  *zap.Logger
	storage storage.Storage
	cfg     *config.Config
}

func NewService(logger *zap.Logger, cfg *config.Config, repo Repository, storage storage.Storage) Service {
	return &serviceImpl{
		repo,
		logger,
		storage,
		cfg,
	}
}

func (s *serviceImpl) GetAllUsers(users *[]models.Users) *apperror.AppError {
	err := s.repo.GetAllUsers(users)

	if err != nil {
		s.logger.Error("Could not get all users", zap.Error(err))
		return apperror.
			New(apperror.InternalServerError).
			Describe("Could not get all users. Please try again later.")
	}

	return nil
}

func (s *serviceImpl) GetUserById(user *models.Users, userId string) *apperror.AppError {
	if !utils.IsValidUUID(userId) {
		return apperror.
			New(apperror.InvalidUserId).
			Describe("Invalid user id")
	}

	err := s.repo.GetUserById(user, userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return apperror.
			New(apperror.UserNotFound).
			Describe("Could not find the specified user")
	} else if err != nil {
		s.logger.Error("Could not get user by id", zap.String("id", userId), zap.Error(err))
		return apperror.
			New(apperror.InternalServerError).
			Describe("Could not get user information. Please try again later.")
	}

	return nil
}

func (s *serviceImpl) Register(user *models.Users, session models.Session) *apperror.AppError {
	var countEmail int64
	if s.repo.CountEmail(&countEmail, user.Email) != nil {
		return apperror.
			New(apperror.InternalServerError).
			Describe("Could not get all emails")
	}

	if countEmail > 0 {
		return apperror.
			New(apperror.EmailAlreadyExists).
			Describe("Email already exists")
	}

	var countPhoneNumber int64
	if s.repo.CountPhoneNumber(&countPhoneNumber, user.PhoneNumber) != nil {
		return apperror.
			New(apperror.InternalServerError).
			Describe("Could not get all phone numbers")
	}

	if countPhoneNumber > 0 {
		return apperror.
			New(apperror.PhoneNumberAlreadyExists).
			Describe("Phone number already exists")
	}

	switch session.RegisteredType {
	case models.EMAIL:
		if !utils.IsValidEmail(user.Email) {
			return apperror.
				New(apperror.InvalidEmail).
				Describe("Invalid email format")
		}

		if !utils.IsValidPassword(user.Password) {
			return apperror.
				New(apperror.InvalidPassword).
				Describe("Password should longer than 8 characters and contain alphabet and numeric characters")
		}

		hashedPassword, hashErr := utils.HashPassword(user.Password)
		if hashErr != nil {
			s.logger.Error("Could not create user", zap.Error(hashErr))
			return apperror.
				New(apperror.InternalServerError).
				Describe("Could not create user. Please try again later")
		}
		user.Password = string(hashedPassword)

	case models.GOOGLE:
		user.Email = session.Email
		user.Password = ""
	}

	user.RegisteredType = session.RegisteredType

	err := s.repo.CreateUser(user)
	if err != nil {
		s.logger.Error("Could not create user", zap.Error(err))
		return apperror.
			New(apperror.InternalServerError).
			Describe("Could not create user. Please try again later")
	}

	return nil
}

func (s *serviceImpl) UpdateUser(user *models.Users, userId string) *apperror.AppError {
	if !utils.IsValidUUID(userId) {
		return apperror.
			New(apperror.InvalidUserId).
			Describe("Invalid user id")
	}

	err := s.repo.UpdateUser(user, userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return apperror.
			New(apperror.UserNotFound).
			Describe("Could not find the specified user")
	} else if err != nil {
		s.logger.Error("Could not update user info", zap.String("id", userId), zap.Error(err))
		return apperror.
			New(apperror.InternalServerError).
			Describe("Could not update user information. Please try again later")
	}

	return nil
}

func (s *serviceImpl) DeleteUser(userId string) *apperror.AppError {
	if !utils.IsValidUUID(userId) {
		return apperror.
			New(apperror.InvalidUserId).
			Describe("Invalid user id")
	}

	err := s.repo.DeleteUser(userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return apperror.
			New(apperror.UserNotFound).
			Describe("Could not find specified user")
	} else if err != nil {
		s.logger.Error("Could not delete user", zap.String("id", userId), zap.Error(err))
		return apperror.
			New(apperror.InternalServerError).
			Describe("Could not delete user. Please try again later")
	}

	return nil
}

func (s *serviceImpl) GetUserByEmail(user *models.Users, email string) *apperror.AppError {
	// Actaully, this shouldn't trigger unless data in database is somehow fucked
	if !utils.IsValidEmail(email) {
		s.logger.Error("Invalid email format", zap.String("email", email))
		return apperror.
			New(apperror.InternalServerError).
			Describe("Invalid email format. Try re-logging in")
	}

	// Same here
	err := s.repo.GetUserByEmail(user, email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		s.logger.Error("Could not find specified user", zap.String("email", email), zap.Error(err))
		return apperror.
			New(apperror.InternalServerError).
			Describe("Could not find specified user. Try re-logging in")
	} else if err != nil {
		s.logger.Error("Could not get current user info", zap.String("email", email), zap.Error(err))
		return apperror.
			New(apperror.InternalServerError).
			Describe("Could not get user information. Please try again later.")
	}

	return nil
}

func (s *serviceImpl) UploadProfileImage(userId uuid.UUID, filename string, file io.Reader) (string, *apperror.AppError) {
	ext := filepath.Ext(filename)
	valid := false
	for _, allowExt := range s.cfg.AllowImageExtensions {
		if allowExt == strings.ToLower(ext[1:]) {
			valid = true
			break
		}
	}

	if !valid {
		return "", apperror.
			New(apperror.InvalidProfileImageExtension).
			Describe(fmt.Sprintf("App does not support %v extension", ext))
	}

	url, err := s.storage.Upload(fmt.Sprintf("profiles/%v%v", userId.String(), ext), file)
	if err != nil {
		return "", apperror.
			New(apperror.InternalServerError).
			Describe("Could not upload profile image")
	}

	return url, nil
}
