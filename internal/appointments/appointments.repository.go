package appointments

import (
	"github.com/brain-flowing-company/pprp-backend/internal/models"
	"gorm.io/gorm"
)

type Repository interface {
	GetAllAppointments(*[]models.Appointments) error
	GetAppointmentsById(*models.Appointments, string) error
	GetAppointmentsByOwnerId(*[]models.Appointments, string) error
	GetAppointmentsByDwellerId(*[]models.Appointments, string) error
	CreateAppointments(*[]models.Appointments) error
	UpdateAppointments(*models.Appointments, string) error
	DeleteAppointments(string) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		db,
	}
}

func (repo *repositoryImpl) GetAllAppointments(results *[]models.Appointments) error {
	return repo.db.Model(&models.Appointments{}).
		Find(results).Error
}

func (repo *repositoryImpl) GetAppointmentsById(result *models.Appointments, id string) error {
	return repo.db.Model(&models.Appointments{}).
		First(result, "appointment_id = ?", id).Error
}

func (repo *repositoryImpl) GetAppointmentsByOwnerId(result *[]models.Appointments, id string) error {
	return repo.db.Model(&models.Appointments{}).
		First(result, "owner_user_id = ?", id).Error
}

func (repo *repositoryImpl) GetAppointmentsByDwellerId(result *[]models.Appointments, id string) error {
	return repo.db.Model(&models.Appointments{}).
		First(result, "dweller_user_id = ?", id).Error
}

func (repo *repositoryImpl) CreateAppointments(apps *[]models.Appointments) error {
	return repo.db.Model(&models.Appointments{}).
		CreateInBatches(apps, len(*apps)).Error
}

func (repo *repositoryImpl) UpdateAppointments(app *models.Appointments, id string) error {
	return repo.db.Model(&models.Appointments{}).
		Where("appointment_id = ?", id).
		Updates(app).Error
}

func (repo *repositoryImpl) DeleteAppointments(id string) error {
	return repo.db.Model(&models.Appointments{}).
		Delete(&models.Appointments{}, "appointment_id = ?", id).Error
}
