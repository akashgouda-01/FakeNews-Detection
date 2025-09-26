// controllers/auth_controller.go
package controllers

import (
	"context"
	"net/http"
	"os"
	"time"

	"mediflow/backend/internal/config"
	"mediflow/backend/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// REMOVED: var userCollection = config.GetCollection("users")

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the collection handle inside the function
		var userCollection = config.GetCollection("users")
		var user models.User
		
		// 1. Bind the incoming JSON to the user model
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 2. Check if a user with this email already exists
		count, err := userCollection.CountDocuments(context.TODO(), bson.M{"email": user.Email})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking for user"})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User with this email already exists"})
			return
		}

		// 3. Hash the password for security
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		user.Password = string(hashedPassword)
		user.ID = primitive.NewObjectID()

		// 4. Insert the new user into the database
		_, err = userCollection.InsertOne(context.TODO(), user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the collection handle inside the function
		var userCollection = config.GetCollection("users")
		var loginDetails struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		// 1. Bind the incoming JSON to the login details struct
		if err := c.BindJSON(&loginDetails); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var foundUser models.User
		// 2. Find the user by email in the database
		err := userCollection.FindOne(context.TODO(), bson.M{"email": loginDetails.Email}).Decode(&foundUser)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		// 3. Compare the provided password with the stored hash
		err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(loginDetails.Password))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		// 4. Generate a JWT if the password is correct
		claims := jwt.MapClaims{
			"userID": foundUser.ID,
			"role":   foundUser.Role,
			"exp":    time.Now().Add(time.Hour * 24).Unix(), // Token is valid for 24 hours
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}

		// 5. Send the token back to the client
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	}
}