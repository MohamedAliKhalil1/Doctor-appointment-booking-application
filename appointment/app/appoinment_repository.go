package app

import "github.com/doctorBooking/appointment/model"

type AppointmentRepository interface {
	SaveAppointment(appt *model.Appointment) (*model.Appointment, error)
}
