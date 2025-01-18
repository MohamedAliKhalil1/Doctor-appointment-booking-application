package http

import (
	"encoding/json"
	"github.com/doctorBooking/appointment/app"
	"github.com/doctorBooking/appointment/model"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type AppointmentHandler struct {
	appointmentUseCase *app.AppointmentReserveUseCase
}

func NewAppointmentHandler(appointmentUseCase *app.AppointmentReserveUseCase) *AppointmentHandler {
	return &AppointmentHandler{appointmentUseCase}
}

func (h *AppointmentHandler) CreateAppointment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var appointment model.Appointment
	err = json.NewDecoder(r.Body).Decode(&appointment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	appointment.SlotID = id
	av, err := h.appointmentUseCase.CreateAppointment(&appointment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(av)
	if err != nil {
		return
	}
}
