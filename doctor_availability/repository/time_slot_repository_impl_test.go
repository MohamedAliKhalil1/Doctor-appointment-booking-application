package repository

import (
	"github.com/doctorBooking/doctor_availability/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func createTimeSlot() *model.TimeSlot {
	return &model.TimeSlot{
		ID:         uuid.New(),
		Time:       time.Now(),
		DoctorID:   uuid.New(),
		DoctorName: "Test Doctor Name",
		IsReserved: false,
		Cost:       100.0,
	}
}

func TestNewTimeSlotRepositoryImpl(t *testing.T) {
	got := NewTimeSlotRepositoryImpl()
	want := &TimeSlotRepositoryImpl{
		timeSlots: make(map[uuid.UUID]*model.TimeSlot),
	}
	assert.Equal(t, want.timeSlots, got.timeSlots, "Expected timeSlots maps to be equal")
	assert.Empty(t, got.timeSlots, "Expected timeSlots maps to be empty")
	assert.IsTypef(t, &sync.RWMutex{}, got.mutex, "Expected the mutex to be of type *sync.RWMutex")
}

func TestTimeSlotRepositoryImpl_AddTimeSlot(t *testing.T) {
	tests := []struct {
		name          string
		timeSlot      *model.TimeSlot
		expectedError bool
	}{
		{
			name:          "add a valid time slot",
			timeSlot:      createTimeSlot(),
			expectedError: false,
		},
		{
			name:          "add nil time slot",
			timeSlot:      nil,
			expectedError: true,
		},
		{
			name: "add a time slot with nil ID",
			timeSlot: func() *model.TimeSlot {
				ts := createTimeSlot()
				ts.ID = uuid.Nil
				return ts
			}(),
			expectedError: false,
		},
		{
			name: "add a time slot with invalid cost",
			timeSlot: func() *model.TimeSlot {
				ts := createTimeSlot()
				ts.Cost = -20
				return ts
			}(),
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewTimeSlotRepositoryImpl()
			err := repo.AddTimeSlot(tt.timeSlot)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestTimeSlotRepositoryImpl_ReserveTimeSlot(t *testing.T) {
	repo := NewTimeSlotRepositoryImpl()
	timeSlot := createTimeSlot()
	_ = repo.AddTimeSlot(timeSlot)
	tests := []struct {
		name          string
		timeSlotId    uuid.UUID
		expectedError bool
	}{
		{
			name:          "reserve a valid time slot",
			timeSlotId:    timeSlot.ID,
			expectedError: false,
		},
		{
			name:          "reserve nil time slot ID",
			timeSlotId:    uuid.Nil,
			expectedError: true,
		},
		{
			name: "reserve reserved time slot",
			timeSlotId: func() uuid.UUID {
				ts := createTimeSlot()
				ts.IsReserved = true
				_ = repo.AddTimeSlot(ts)
				return ts.ID
			}(),
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slot, err := repo.ReserveTimeSlot(tt.timeSlotId)
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, slot)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, slot)
			}
		})
	}
}

func TestTimeSlotRepositoryImpl_ListTimeSlots(t *testing.T) {
	repo := NewTimeSlotRepositoryImpl()
	slotsNum := 10
	for i := 0; i < slotsNum; i++ {
		timeSlot := createTimeSlot()
		_ = repo.AddTimeSlot(timeSlot)
	}

	want := slotsNum
	got := len(repo.ListTimeSlots())
	assert.Equal(t, want, got)
}
