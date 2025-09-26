package controllers

import (
	"context"
	"net/http"
	"time"

	"mediflow/backend/internal/config"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// REMOVED: var bedCollection = config.GetCollection("beds")

func GetAllBeds() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the collection handle INSIDE the function
		var bedCollection = config.GetCollection("beds")
		
		cursor, err := bedCollection.Find(context.TODO(), bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve beds"})
			return
		}
		defer cursor.Close(context.TODO())

		var beds []bson.M
		if err = cursor.All(context.TODO(), &beds); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding bed data"})
			return
		}

		c.JSON(http.StatusOK, beds)
	}
}

func AllocateBed() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the collection handle INSIDE the function
		var bedCollection = config.GetCollection("beds")
		
		bedID, err := primitive.ObjectIDFromHex(c.Param("bedId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bed ID"})
			return
		}

		var allocationDetails struct {
			PatientID string `json:"patientId"`
		}
		if err := c.BindJSON(&allocationDetails); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		update := bson.M{
			"$set": bson.M{
				"status":              "Occupied",
				"occupiedByPatientId": allocationDetails.PatientID,
				"timestamps.allocatedDate": time.Now(),
			},
		}

		result, err := bedCollection.UpdateOne(context.TODO(), bson.M{"_id": bedID, "status": "Available"}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to allocate bed"})
			return
		}

		if result.MatchedCount == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Bed not found or is already occupied"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Bed allocated successfully"})
	}
}

func DischargeBed() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the collection handle INSIDE the function
		var bedCollection = config.GetCollection("beds")
		
		bedID, err := primitive.ObjectIDFromHex(c.Param("bedId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bed ID"})
			return
		}
		
		update := bson.M{
			"$set": bson.M{
				"status":                 "Cleaning",
				"occupiedByPatientId":    "",
				"timestamps.dischargedDate": time.Now(),
			},
		}

		_, err = bedCollection.UpdateOne(context.TODO(), bson.M{"_id": bedID}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to discharge bed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Bed discharged successfully, status set to Cleaning"})
	}
}