package app

import (
	"github.com/doctorBooking/doctor_availability/model"
	"github.com/google/uuid"
)

type TimeSlotRepository interface {
	ListTimeSlots() []*model.TimeSlot
	AddTimeSlot(timeSlot *model.TimeSlot) error
	ReserveTimeSlot(slotId uuid.UUID) (*model.TimeSlot, error)
}
