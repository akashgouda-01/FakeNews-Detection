package controllers

import (
	"context"
	"net/http"
	"time"

	"mediflow/backend/internal/config"
	"mediflow/backend/internal/models"

	"github.comcom/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// REMOVED: var patientCollection = config.GetCollection("patients")

func AdmitPatient() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the collection handle INSIDE the function
		var patientCollection = config.GetCollection("patients")
		var patient models.Patient
		
		if err := c.BindJSON(&patient); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		patient.ID = primitive.NewObjectID()
		if len(patient.Admissions) > 0 {
			patient.Admissions[0].AdmissionID = "ADM-" + primitive.NewObjectID().Hex()
			patient.Admissions[0].AdmissionDate = time.Now()
		}

		_, err := patientCollection.InsertOne(context.TODO(), patient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to admit patient"})
			return
		}

		c.JSON(http.StatusCreated, patient)
	}
}

func GetAllPatients() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the collection handle INSIDE the function
		var patientCollection = config.GetCollection("patients")
		var patients []models.Patient
		
		cursor, err := patientCollection.Find(context.TODO(), bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve patients"})
			return
		}
		defer cursor.Close(context.TODO())

		if err = cursor.All(context.TODO(), &patients); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding patient data"})
			return
		}

		c.JSON(http.StatusOK, patients)
	}
}