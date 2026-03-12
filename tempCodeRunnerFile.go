package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {

	// MongoDB connect
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB not connected")
	}

	fmt.Println("MongoDB Connected ✅")

	collection = client.Database("devops_security").Collection("users")

	router := gin.Default()

	// Register API
	router.POST("/register", func(c *gin.Context) {

		var user User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		_, err := collection.InsertOne(context.TODO(), user)

		if err != nil {
			c.JSON(500, gin.H{"error": "User not saved"})
			return
		}

		c.JSON(200, gin.H{"message": "User Registered"})
	})

	// Login API
	router.POST("/login", func(c *gin.Context) {

		var user User
		var dbUser User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		err := collection.FindOne(context.TODO(),
			map[string]string{"username": user.Username}).Decode(&dbUser)

		if err != nil {
			c.JSON(401, gin.H{"error": "User not found"})
			return
		}

		if user.Password != dbUser.Password {
			c.JSON(401, gin.H{"error": "Wrong password"})
			return
		}

		c.JSON(200, gin.H{"message": "Login Successful"})
	})

	router.Run(":9092")
}