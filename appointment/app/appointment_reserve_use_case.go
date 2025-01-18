package app

import (
	"errors"
	"github.com/doctorBooking/appointment/model"
	"github.com/doctorBooking/appointment_confirmation"
	"github.com/doctorBooking/doctor_availability/app"
	"github.com/google/uuid"
)

type AppointmentReserveUseCase struct {
	appointmentRepository AppointmentRepository
	timeSlotService       app.ItimeSlotService
	notify                appointment_confirmation.AppointmentConfirmationService
}

func NewAppointmentReserveUseCase(repo AppointmentRepository, timeSlotService app.ItimeSlotService, notify appointment_confirmation.AppointmentConfirmationService) *AppointmentReserveUseCase {
	return &AppointmentReserveUseCase{
		appointmentRepository: repo,
		timeSlotService:       timeSlotService,
		notify:                notify,
	}
}

func (apptuc *AppointmentReserveUseCase) CreateAppointment(appt *model.Appointment) (*model.Appointment, error) {
	if appt == nil {
		return nil, errors.New("appointment cannot be nil")
	}

	if appt.ID == uuid.Nil {
		appt.ID = uuid.New()
	}

	slot, err := apptuc.timeSlotService.ReserveTimeSlot(appt.SlotID)
	if err != nil {
		return nil, err
	}

	createdAppointment, err := apptuc.appointmentRepository.SaveAppointment(appt)
	if err != nil {
		return nil, err
	}
	err = apptuc.notify.SendConfirmation(createdAppointment.PatientName, slot.DoctorName, slot.Time)
	if err != nil {
		return nil, err
	}
	return createdAppointment, nil
}
