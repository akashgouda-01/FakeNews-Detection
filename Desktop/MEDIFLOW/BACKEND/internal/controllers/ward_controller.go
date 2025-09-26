package controllers

import (
	"context"
	"fmt"
	"net/http"

	"mediflow/backend/internal/config"
	"mediflow/backend/internal/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// REMOVED: var wardCollection = config.GetCollection("wards")

func GetAllWards() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the collection handle INSIDE the function
		var wardCollection = config.GetCollection("wards")
		var wards []models.Ward
		
		cursor, err := wardCollection.Find(context.TODO(), bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve wards"})
			return
		}
		defer cursor.Close(context.TODO())

		if err = cursor.All(context.TODO(), &wards); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding ward data"})
			return
		}

		c.JSON(http.StatusOK, wards)
	}
}

func SeedWards() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the collection handle INSIDE the function
		var wardCollection = config.GetCollection("wards")
		
		wardNames := []string{
			"ICU", "Emergency", "Surgery", "Maternity", "Pediatrics",
			"Psychiatry", "Rehabilitation", "Oncology", "Cardiology", "Neurology",
			"Orthopedics", "ENT", "Dermatology", "Ophthalmology", "Urology",
			"Gynecology", "Neonatal", "Burns Unit", "Isolation", "Day Care",
		}

		var wardsToInsert []interface{}
		for _, name := range wardNames {
			count, err := wardCollection.CountDocuments(context.TODO(), bson.M{"name": name})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check for existing wards"})
				return
			}
			if count > 0 {
				continue 
			}

			newWard := models.Ward{
				ID:           primitive.NewObjectID(),
				Name:         name,
				TotalBeds:    20,
				OccupiedBeds: 0,
				PatientIDs:   []string{},
			}
			wardsToInsert = append(wardsToInsert, newWard)
		}

		if len(wardsToInsert) == 0 {
			c.JSON(http.StatusOK, gin.H{"message": "All wards already exist in the database."})
			return
		}

		_, err := wardCollection.InsertMany(context.TODO(), wardsToInsert)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to seed wards"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("Successfully seeded %d new wards.", len(wardsToInsert))})
	}
}