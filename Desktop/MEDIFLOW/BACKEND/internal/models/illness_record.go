package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Location stores geographic coordinates for the heatmap.
type Location struct {
	Lat  float64 `bson:"lat" json:"lat"`
	Long float64 `bson:"long" json:"long"`
}

// IllnessRecord tracks a single diagnosed illness for epidemiological analysis.
type IllnessRecord struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	PatientID        string             `bson:"patientId" json:"patientId"`               // [cite: 192]
	Diagnosis        string             `bson:"diagnosis" json:"diagnosis"`               // [cite: 193]
	DateDiagnosed    time.Time          `bson:"dateDiagnosed" json:"dateDiagnosed"`       // [cite: 194]
	HospitalLocation Location           `bson:"hospitalLocation" json:"hospitalLocation"` // [cite: 195]
}