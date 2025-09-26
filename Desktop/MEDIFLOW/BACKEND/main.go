package main

import (
	"mediflow/backend/internal/config"
	"mediflow/backend/internal/controllers"
	//	"mediflow/backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Initialize the Gin router
	router := gin.Default()

	// 2. Connect to the database
	config.ConnectDB()

	// 3. Define Route Groups
	
	// Public routes for dashboard login/registration (no token required)
	authRoutes := router.Group("/api/auth")
	{
		authRoutes.POST("/register", controllers.Register())
		authRoutes.POST("/login", controllers.Login())
	}

	// Protected API routes (require a valid JWT)
	api := router.Group("/api")
	// api.Use(middleware.AuthMiddleware())
	{
		// --- Patient Routes ---
		// Corresponds to the "Patient Admissions Form" [cite: 104]
		api.POST("/patients/admit", controllers.AdmitPatient())
		api.GET("/patients", controllers.GetAllPatients())

		// --- Bed Routes ---
		// For the "Bed Allocation & Management Dashboard" [cite: 91]
		api.GET("/beds", controllers.GetAllBeds())
		api.PUT("/beds/:bedId/allocate", controllers.AllocateBed())   // Corresponds to "Patient to Bed Allocation" [cite: 98]
		api.PUT("/beds/:bedId/discharge", controllers.DischargeBed()) // Corresponds to "Discharge & Release Bed" [cite: 99]

		// --- Ward Routes ---
		api.GET("/wards", controllers.GetAllWards())
		api.POST("/wards/seed", controllers.SeedWards()) // Use this once to populate wards

		// --- Illness Routes (Placeholder) ---
		// For the "Illness Insights Dashboard" [cite: 106]
		// api.GET("/illness-records", controllers.GetIllnessRecords())

		// --- Blood Bank Routes (Placeholder) ---
		// For the "Blood Bank Management Dashboard" [cite: 110]
		// api.GET("/blood-inventory", controllers.GetBloodInventory())
	}

	// 4. Start the server
	router.Run("localhost:8080")
}