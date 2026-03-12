package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var jwtKey = []byte("devops_secret_key")

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func generateToken(username string) (string, error) {

	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}

func authMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token required",
			})
			c.Abort()
			return
		}

		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func main() {

	// MongoDB connection
	clientOptions := options.Client().ApplyURI("mongodb://mongodb-service:27017")

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal("MongoDB connection failed")
	}

	fmt.Println("MongoDB Connected ✅")

	collection = client.Database("devops_security").Collection("users")

	router := gin.Default()

	// Register API
	router.POST("/register", func(c *gin.Context) {

		var user User

		if err := c.BindJSON(&user); err != nil {

			c.JSON(400, gin.H{
				"error": "Invalid input",
			})
			return
		}

		_, err := collection.InsertOne(context.Background(), user)

		if err != nil {

			c.JSON(500, gin.H{
				"error": "User not saved",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "User Registered Successfully",
		})
	})

	// Login API
	router.POST("/login", func(c *gin.Context) {

		var user User
		var dbUser User

		if err := c.BindJSON(&user); err != nil {

			c.JSON(400, gin.H{
				"error": "Invalid input",
			})
			return
		}

		err := collection.FindOne(context.Background(),
			map[string]string{"username": user.Username}).Decode(&dbUser)

		if err != nil {

			c.JSON(401, gin.H{
				"error": "User not found",
			})
			return
		}

		if user.Password != dbUser.Password {

			c.JSON(401, gin.H{
				"error": "Wrong password",
			})
			return
		}

		token, err := generateToken(user.Username)

		if err != nil {

			c.JSON(500, gin.H{
				"error": "Token generation failed",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Login Successful",
			"token":   token,
		})
	})

	// Protected API
	router.GET("/secure-data", authMiddleware(), func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "This is protected DevOps security data",
		})
	})

	router.Run(":9092")
}
