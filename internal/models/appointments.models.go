package models

import (
	"time"

	"github.com/brain-flowing-company/pprp-backend/internal/enums"
	"github.com/google/uuid"
)

type Appointments struct {
	AppointmentId    uuid.UUID               `json:"appointment_id"   example:"123e4567-e89b-12d3-a456-426614174000"`
	PropertyId       uuid.UUID               `json:"property_id"      example:"123e4567-e89b-12d3-a456-426614174000"`
	OwnerUserId      uuid.UUID               `json:"owner_user_id"    example:"123e4567-e89b-12d3-a456-426614174000"`
	DwellerUserId    uuid.UUID               `json:"dweller_user_id"  example:"123e4567-e89b-12d3-a456-426614174000"`
	Status           enums.AppointmentStatus `json:"status"           example:"PENDING"`
	AppointmentDate  time.Time               `json:"appointment_date" example:"2024-02-18T11:00:00Z"`
	Note             string                  `json:"note"             example:"This is a note"`
	CancelledMessage string                  `json:"cancelled_message" example:"This is a cancelled message"`
	CommonModels
}

type CreatingAppointments struct {
	PropertyId       uuid.UUID `json:"property_id"       example:"123e4567-e89b-12d3-a456-426614174000"`
	OwnerUserId      uuid.UUID `json:"owner_user_id"     example:"123e4567-e89b-12d3-a456-426614174000"`
	DwellerUserId    uuid.UUID `json:"dweller_user_id"   example:"123e4567-e89b-12d3-a456-426614174000"`
	AppointmentDate  time.Time `json:"appointment_dates" example:"2024-02-18T11:00:00Z"`
	Note             string    `json:"note"             example:"This is a note"`
	CancelledMessage string    `json:"cancelled_message" example:"This is a cancelled message"`
}

func (a Appointments) TableName() string {
	return "appointments"
}

type UpdatingAppointmentStatus struct {
	Status enums.AppointmentStatus `json:"status"`
}
