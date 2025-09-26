package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Ward represents a hospital ward.
type Ward struct {
    ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Name         string             `bson:"name" json:"name"`
    TotalBeds    int                `bson:"totalBeds" json:"totalBeds"`
    OccupiedBeds int                `bson:"occupiedBeds" json:"occupiedBeds"`
    PatientIDs   []string           `bson:"patientIds" json:"patientIds"`
}