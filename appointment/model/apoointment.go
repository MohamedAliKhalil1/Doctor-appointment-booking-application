package model

import "github.com/google/uuid"

type Appointment struct {
	ID          uuid.UUID `json:"id"`
	SlotID      uuid.UUID `json:"slot_id"`
	PatientID   uuid.UUID `json:"patient_id"`
	PatientName string    `json:"patient_name"`
	ReservedAt  string    `json:"reserved_at"`
}
