package appointment_confirmation

import (
	"errors"
	"log"
	"time"
)

type AppointmentConfirmationService interface {
	SendConfirmation(patientName, doctorName string, appointmentTime time.Time) error
}

type AppointmentConfirmationServiceImpl struct{}

func NewAppointmentConfirmationServiceImpl() *AppointmentConfirmationServiceImpl {
	return &AppointmentConfirmationServiceImpl{}
}

func (apptc AppointmentConfirmationServiceImpl) SendConfirmation(patientName, doctorName string, appointmentTime time.Time) error {
	if patientName == "" || doctorName == "" || appointmentTime.IsZero() {
		return errors.New("missing required parameters")
	}
	log.Printf("Appointment confirmed for %s with Dr. %s at %s", patientName, doctorName, appointmentTime)
	return nil
}
