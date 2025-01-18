package repository

import (
	"fmt"
	"github.com/doctorBooking/doctor_availability/model"
	"github.com/google/uuid"
	"log"
	"sync"
)

type TimeSlotRepositoryImpl struct {
	TimeSlots map[uuid.UUID]*model.TimeSlot
	Mutex     *sync.RWMutex
}

func NewTimeSlotRepositoryImpl() *TimeSlotRepositoryImpl {
	timeSlots := make(map[uuid.UUID]*model.TimeSlot)
	mutex := &sync.RWMutex{}
	return &TimeSlotRepositoryImpl{TimeSlots: timeSlots, Mutex: mutex}
}

func (repo *TimeSlotRepositoryImpl) CopyTimeSlote(timeSlot *model.TimeSlot) (*model.TimeSlot, error) {
	if timeSlot == nil {
		return nil, fmt.Errorf("time slot is nil")
	}
	return &model.TimeSlot{
		ID:         timeSlot.ID,
		Time:       timeSlot.Time,
		DoctorID:   timeSlot.DoctorID,
		DoctorName: timeSlot.DoctorName,
		Cost:       timeSlot.Cost,
		IsReserved: timeSlot.IsReserved,
	}, nil
}

func (repo *TimeSlotRepositoryImpl) ListTimeSlots() []*model.TimeSlot {
	repo.Mutex.RLock()
	defer repo.Mutex.RUnlock()

	timeSlots := make([]*model.TimeSlot, 0, len(repo.TimeSlots))
	for _, slot := range repo.TimeSlots {
		slotCopy, err := repo.CopyTimeSlote(slot)
		if err != nil {
			log.Println(err)
		}
		timeSlots = append(timeSlots, slotCopy)
	}
	return timeSlots
}

func (repo *TimeSlotRepositoryImpl) AddTimeSlot(timeSlot *model.TimeSlot) error {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	if timeSlot == nil {
		return fmt.Errorf("time slot is nil")
	}

	if timeSlot.ID == uuid.Nil {
		timeSlot.ID = uuid.New()
	}

	if _, exists := repo.TimeSlots[timeSlot.ID]; exists {
		return fmt.Errorf("time slot with ID %s already exists", timeSlot.ID)
	}

	if timeSlot.Cost < 0 {
		return fmt.Errorf("time slot cost is negative")
	}
	repo.TimeSlots[timeSlot.ID] = timeSlot
	return nil
}

func (repo *TimeSlotRepositoryImpl) ReserveTimeSlot(slotId uuid.UUID) (
	*model.TimeSlot,
	error,
) {
	repo.Mutex.Lock()
	defer repo.Mutex.Unlock()

	if _, exists := repo.TimeSlots[slotId]; !exists {
		return nil, fmt.Errorf("time slot with ID %s does not exist", slotId)
	}
	slot := repo.TimeSlots[slotId]
	if slot.IsReserved {
		return nil, fmt.Errorf("time slot with ID %s is already reserved", slotId)
	}
	slot.IsReserved = true
	return repo.CopyTimeSlote(slot)
}

func (repo *TimeSlotRepositoryImpl) GetTimeSlot(slotId uuid.UUID) (*model.TimeSlot, error) {
	repo.Mutex.RLock()
	defer repo.Mutex.RUnlock()
	if _, exists := repo.TimeSlots[slotId]; !exists {
		return nil, fmt.Errorf("time slot with ID %s does not exist", slotId)
	}
	return repo.CopyTimeSlote(repo.TimeSlots[slotId])
}
