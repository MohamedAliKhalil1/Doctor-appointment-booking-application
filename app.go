package main

import (
	"github.com/doctorBooking/doctor_availability/app"
	httpav "github.com/doctorBooking/doctor_availability/http"
	"github.com/doctorBooking/doctor_availability/repository"
	"github.com/gorilla/mux"
	http "net/http"
)

func main() {
	timeSlotRepo := repository.NewTimeSlotRepositoryImpl()
	timeSlotService := app.NewTimeSlotService(timeSlotRepo)
	timeSlotHandler := httpav.NewTimeSlotHandler(timeSlotService)

	r := mux.NewRouter()
	r.HandleFunc("/slots", timeSlotHandler.ListTimeSlots).Methods("GET")
	r.HandleFunc("/slots/{id}", timeSlotHandler.GetTimeSlot).Methods("GET")
	r.HandleFunc("/slots", timeSlotHandler.CreateTimeSlot).Methods("POST")
	r.HandleFunc("/slots/{id}/reserve", timeSlotHandler.ReserveTimeSlot).Methods("POST")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
