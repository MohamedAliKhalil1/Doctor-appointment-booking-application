package app

import (
	"errors"
	"github.com/doctorBooking/appointment/model"
	"github.com/google/uuid"
)

type AppointmentReserveUseCase struct {
	repo AppointmentRepository
}

func NewAppointmentReserveUseCase(repo AppointmentRepository) *AppointmentReserveUseCase {
	return &AppointmentReserveUseCase{repo: repo}
}

func (apptuc *AppointmentReserveUseCase) CreateAppointment(appt *model.Appointment) error {
	if appt == nil {
		return errors.New("appointment cannot be nil")
	}

	if appt.ID == uuid.Nil {
		appt.ID = uuid.New()
	}

	return apptuc.repo.SaveAppointment(appt)
}
