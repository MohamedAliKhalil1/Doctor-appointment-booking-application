package model

import (
	"github.com/google/uuid"
	"time"
)

type TimeSlot struct {
	ID         uuid.UUID `json:"id"`
	Time       time.Time `json:"time"`
	DoctorID   uuid.UUID `json:"doctor_id"`
	DoctorName string    `json:"doctor_name"`
	IsReserved bool      `json:"is_reserved"`
	Cost       float64   `json:"cost"`
}
