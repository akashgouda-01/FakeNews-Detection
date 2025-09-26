package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Admission represents a patient's admission record.
type Admission struct {
	AdmissionID      string    `bson:"admissionId" json:"admissionId"`            // [cite: 138]
	AdmissionDate    time.Time `bson:"admissionDate" json:"admissionDate"`        // [cite: 139]
	Ward             string    `bson:"ward" json:"ward"`                          // [cite: 140]
	ReasonForAdmitting string    `bson:"reasonForAdmitting" json:"reasonForAdmitting"` // [cite: 141]
	ConsultingDoctor string    `bson:"consultingDoctor" json:"consultingDoctor"`  // [cite: 142]
	DischargeDate    time.Time `bson:"dischargeDate,omitempty" json:"dischargeDate"` // [cite: 143]
}

// PersonalDetails contains the demographic information of a patient.
type PersonalDetails struct {
	Name          string    `bson:"name" json:"name"`                      // [cite: 128]
	Age           int       `bson:"age" json:"age"`                        // [cite: 129]
	DOB           time.Time `bson:"dob" json:"dob"`                        // [cite: 130]
	BloodGroup    string    `bson:"bloodGroup" json:"bloodGroup"`          // [cite: 131]
	Gender        string    `bson:"gender" json:"gender"`                  // [cite: 132]
	Guardian      string    `bson:"guardian" json:"guardian"`              // [cite: 133]
	ContactNumber string    `bson:"contactNumber" json:"contactNumber"`    // [cite: 134]
	Address       string    `bson:"address" json:"address"`                // [cite: 135]
}

// Patient represents the main document in the 'patients' collection.
type Patient struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`   // [cite: 126]
	PersonalDetails PersonalDetails    `bson:"personalDetails" json:"personalDetails"` // [cite: 127]
	Admissions      []Admission        `bson:"admissions" json:"admissions"` // [cite: 136]
}