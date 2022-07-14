package main

import (
	// Standard library packages

	// Third party packages
	"GolangMongo/controllers"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2"
)

func main() {

	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	ucPost := controllers.NewPostController(getSession())

	// Get a user resource
	router := gin.Default()
	router.GET("/users", uc.UsersList)
	router.GET("/users/:id", uc.GetUser)
	router.DELETE("/users/:id", uc.RemoveUser)
	router.POST("/users", uc.CreateUser)
	router.PUT("/users/:id", uc.UpdateUser)

	router.GET("/posts", ucPost.PostsList)
	router.GET("/posts/:id", ucPost.GetPost)
	router.GET("/posts_by_slug/:id", ucPost.GetPostBySlug)
	router.DELETE("/posts/:id", ucPost.RemovePost)
	router.POST("/posts", ucPost.CreatePost)
	router.PUT("/posts/:id", ucPost.UpdatePost)

	router.Run(":8585")
}

// getSession creates a new mongo session and panics if connection error occurs
func getSession() *mgo.Session {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	hosts := os.Getenv("hosts")
	username := os.Getenv("username")
	password := os.Getenv("password")

	// Connect to our local mongo
	info := &mgo.DialInfo{
		Addrs:   []string{hosts},
		Timeout: 60 * time.Second,
		// Database: database,
		Username: username,
		Password: password,
	}

	s, err := mgo.DialWithInfo(info)

	// s, err := mgo.Dial("mongodb://<dbuser>:<dbpassword>@ds041154.mongolab.com:41154/location")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	// Deliver session
	return s
}
