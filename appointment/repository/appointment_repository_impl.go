package repository

import (
	"errors"
	"github.com/doctorBooking/appointment/model"
	"github.com/google/uuid"
	"sync"
)

type AppointmentRepositoryImpl struct {
	appointments map[uuid.UUID]*model.Appointment
	mutex        *sync.Mutex
}

func NewAppointmentRepositoryImpl() *AppointmentRepositoryImpl {
	return &AppointmentRepositoryImpl{appointments: make(map[uuid.UUID]*model.Appointment), mutex: &sync.Mutex{}}
}

func (a *AppointmentRepositoryImpl) SaveAppointment(appt *model.Appointment) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if appt == nil {
		return errors.New("appointment is nil")
	}

	if _, ok := a.appointments[appt.ID]; ok {
		return errors.New("appointment already exists")
	}

	a.appointments[appt.ID] = appt
	return nil
}
