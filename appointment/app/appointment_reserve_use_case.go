package app

import (
	"errors"
	"github.com/doctorBooking/appointment/model"
	"github.com/doctorBooking/doctor_availability/app"
	"github.com/google/uuid"
)

type AppointmentReserveUseCase struct {
	appointmentRepository AppointmentRepository
	timeSlotService       app.ItimeSlotService
}

func NewAppointmentReserveUseCase(repo AppointmentRepository, timeSlotService app.ItimeSlotService) *AppointmentReserveUseCase {
	return &AppointmentReserveUseCase{appointmentRepository: repo, timeSlotService: timeSlotService}
}

func (apptuc *AppointmentReserveUseCase) CreateAppointment(appt *model.Appointment) (*model.Appointment, error) {
	if appt == nil {
		return nil, errors.New("appointment cannot be nil")
	}

	if appt.ID == uuid.Nil {
		appt.ID = uuid.New()
	}

	_, err := apptuc.timeSlotService.ReserveTimeSlot(appt.SlotID)
	if err != nil {
		return nil, err
	}

	return apptuc.appointmentRepository.SaveAppointment(appt)
}
