package app

import (
	"errors"
	"github.com/doctorBooking/doctor_availability/model"
	"github.com/google/uuid"
)

type ItimeSlotService interface {
	ListTimeSlots() []*model.TimeSlot
	GetTimeSlot(id uuid.UUID) (*model.TimeSlot, error)
	CreateTimeSlot(slot *model.TimeSlot) (*model.TimeSlot, error)
	ReserveTimeSlot(slotId uuid.UUID) (*model.TimeSlot, error)
}

type TimeSlotService struct {
	timeSlotRepo TimeSlotRepository
}

func NewTimeSlotService(timeSlotRepository TimeSlotRepository) *TimeSlotService {
	return &TimeSlotService{timeSlotRepo: timeSlotRepository}
}

func (ts *TimeSlotService) ListTimeSlots() []*model.TimeSlot {
	return ts.timeSlotRepo.ListTimeSlots()
}

func (ts *TimeSlotService) GetTimeSlot(id uuid.UUID) (*model.TimeSlot, error) {
	slot, err := ts.timeSlotRepo.GetTimeSlot(id)
	if err != nil {
		return nil, err
	}
	return slot, nil
}

func (ts *TimeSlotService) CreateTimeSlot(slot *model.TimeSlot) (*model.TimeSlot, error) {
	if slot == nil {
		return nil, errors.New("slot must not be nil")
	}
	if slot.ID == uuid.Nil {
		slot.ID = uuid.New()
	}

	if slot.DoctorID == uuid.Nil {
		return nil, errors.New("slot id must not be nil")
	}

	if slot.Time.IsZero() {
		return nil, errors.New("slot time must not be nil")
	}

	err := ts.timeSlotRepo.AddTimeSlot(slot)
	if err != nil {
		return nil, err
	}
	return slot, nil
}

func (ts *TimeSlotService) ReserveTimeSlot(slotId uuid.UUID) (*model.TimeSlot, error) {
	if slotId == uuid.Nil {
		return nil, errors.New("slot id must not be nil")
	}
	return ts.timeSlotRepo.ReserveTimeSlot(slotId)
}
