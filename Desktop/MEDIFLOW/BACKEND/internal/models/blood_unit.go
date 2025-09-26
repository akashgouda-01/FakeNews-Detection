package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Transaction holds the history of blood unit additions or withdrawals.
type Transaction struct {
	TransactionType string    `bson:"transactionType" json:"transactionType"` // [cite: 203]
	UnitsChanged    int       `bson:"unitsChanged" json:"unitsChanged"`       // [cite: 203]
	Timestamp       time.Time `bson:"timestamp" json:"timestamp"`             // [cite: 203]
}

// BloodUnit represents an inventory item in the blood bank.
type BloodUnit struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	BloodType      string             `bson:"bloodType" json:"bloodType"`           // [cite: 200]
	UnitsAvailable int                `bson:"unitsAvailable" json:"unitsAvailable"` // [cite: 201]
	LastUpdated    time.Time          `bson:"lastUpdated" json:"lastUpdated"`       // [cite: 202]
	History        []Transaction      `bson:"history" json:"history"`               // [cite: 203]
}