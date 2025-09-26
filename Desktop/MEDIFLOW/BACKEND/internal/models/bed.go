package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BedTimestamps tracks allocation and discharge times for a bed.
type BedTimestamps struct {
	AllocatedDate  time.Time `bson:"allocatedDate" json:"allocatedDate"`   // [cite: 212]
	DischargedDate time.Time `bson:"dischargedDate" json:"dischargedDate"` // [cite: 213]
}

// Bed represents a single bed in the hospital.
type Bed struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	BedNumber           string             `bson:"bedNumber" json:"bedNumber"`                     // [cite: 206]
	WardName            string             `bson:"wardName" json:"wardName"`                       // [cite: 207]
	Status              string             `bson:"status" json:"status"`                           // [cite: 208]
	OccupiedByPatientID string             `bson:"occupiedByPatientId,omitempty" json:"patientId"` // [cite: 209]
	BedType             string             `bson:"bedType" json:"bedType"`                         // [cite: 210]
	Timestamps          BedTimestamps      `bson:"timestamps" json:"timestamps"`                   // [cite: 211]
}