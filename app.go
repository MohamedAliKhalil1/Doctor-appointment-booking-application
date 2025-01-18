package main

import (
	appointmentUC "github.com/doctorBooking/appointment/app"
	appointmentH "github.com/doctorBooking/appointment/http"
	appointmentRepo "github.com/doctorBooking/appointment/repository"
	"github.com/doctorBooking/doctor_availability/app"
	httpav "github.com/doctorBooking/doctor_availability/http"
	timeSlotRepo "github.com/doctorBooking/doctor_availability/repository"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	timeSlotRepository := timeSlotRepo.NewTimeSlotRepositoryImpl()
	timeSlotService := app.NewTimeSlotService(timeSlotRepository)
	timeSlotHandler := httpav.NewTimeSlotHandler(timeSlotService)

	appointmentRepository := appointmentRepo.NewAppointmentRepositoryImpl()
	appointmentUseCase := appointmentUC.NewAppointmentReserveUseCase(appointmentRepository, timeSlotService)
	appointmentHandler := appointmentH.NewAppointmentHandler(appointmentUseCase)

	r := mux.NewRouter()
	r.HandleFunc("/slots", timeSlotHandler.ListTimeSlots).Methods("GET")
	r.HandleFunc("/slots/{id}", timeSlotHandler.GetTimeSlot).Methods("GET")
	r.HandleFunc("/slots", timeSlotHandler.CreateTimeSlot).Methods("POST")
	r.HandleFunc("/slots/{id}/reserve", appointmentHandler.CreateAppointment).Methods("POST")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
