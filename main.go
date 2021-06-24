package main

import (
	"context"
	"net/http"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

type user struct {
	Title     string
	Email     string
	Password  string
	FirstName []string
	LastName  []string
	Address   []string
	City      []string
	Country   []string
	Code      []string
	Image     string
	Category  string
}

var router *gin.Engine

// var app *firebase.App
// var ctx context.Context
var client *firestore.Client

var usermodel user
var ctx = context.Background()

//init firebase
var config = &firebase.Config{
	StorageBucket: "dot-cms-f3582.appspot.com",
}
var opt = option.WithCredentialsFile("firebase.json")
var app, err = firebase.NewApp(ctx, config, opt)

func main() {

	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	router.Use(CORSMiddleware())

	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")

	initializeRoutes()

	// Start serving the application
	router.Run()
}

func render(c *gin.Context, data gin.H, templateName string) {
	loggedInInterface, _ := c.Get("is_logged_in")
	data["is_logged_in"] = loggedInInterface.(bool)

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}
