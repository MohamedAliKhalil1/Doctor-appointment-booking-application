package http

import (
	"encoding/json"
	"github.com/doctorBooking/appointment/app"
	"github.com/doctorBooking/appointment/model"
	"net/http"
)

type AppointmentHandler struct {
	appointmentUseCase *app.AppointmentReserveUseCase
}

func NewAppointmentHandler(appointmentUseCase *app.AppointmentReserveUseCase) *AppointmentHandler {
	return &AppointmentHandler{appointmentUseCase}
}

func (h *AppointmentHandler) CreateAppointment(w http.ResponseWriter, r *http.Request) {
	var appointment model.Appointment
	err := json.NewDecoder(r.Body).Decode(&appointment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	err = h.appointmentUseCase.CreateAppointment(&appointment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
