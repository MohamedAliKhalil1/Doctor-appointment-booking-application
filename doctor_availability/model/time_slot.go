package model

import (
	"github.com/google/uuid"
	"time"
)

type TimeSlot struct {
	ID         uuid.UUID
	Time       time.Time
	DoctorID   uuid.UUID
	DoctorName string
	IsReserved bool
	Cost       float64
}
