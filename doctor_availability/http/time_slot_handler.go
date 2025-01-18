package http

import (
	"encoding/json"
	"github.com/doctorBooking/doctor_availability/app"
	"github.com/doctorBooking/doctor_availability/model"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type TimeSlotHandler struct {
	service *app.TimeSlotService
}

func NewTimeSlotHandler(service *app.TimeSlotService) *TimeSlotHandler {
	return &TimeSlotHandler{service: service}
}

func (h *TimeSlotHandler) ListTimeSlots(w http.ResponseWriter, r *http.Request) {
	slots := h.service.ListTimeSlots()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(slots)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *TimeSlotHandler) GetTimeSlot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	slot, err := h.service.GetTimeSlot(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(slot)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *TimeSlotHandler) CreateTimeSlot(w http.ResponseWriter, r *http.Request) {
	var slot model.TimeSlot
	if r.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&slot)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.service.CreateTimeSlot(&slot)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(slot)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *TimeSlotHandler) ReserveTimeSlot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	slot, err := h.service.ReserveTimeSlot(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(slot)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
