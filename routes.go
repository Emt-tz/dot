package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var routes = []string{"afrecast.html", "beneficiariesSE.html", "cim.html", "circle.html", "dashboard.html", "digitalbus.html", "dn.html", "ewong.html", "index.html", "jembe.html", "moow.html", "notifications.html", "programs.html", "socialent.html", "survey.html", "tables.html", "tyds.html", "user.html", "explore.html"}

//user logged in model
type user struct {
	Title     string
	Email     string
	FirstName string
	LastName  string
	Address   string
	City      string
	Country   string
	Code      string
	Contact   string
	Image     string
}

var Title = "Manager"
var Email = "peterkelvin16@gmail.com"
var FirstName = "Emmanuel"
var LastName = "Mtera"
var Address = "Mbezi Jogoo"
var City = "Dar es salaam"
var Country = "Tanzania"
var Code = "14128"
var Contact = "0744568803"
var Image = "../assets/img/profile-img01.jpg"

var usermodel = user{Title, Email, FirstName, LastName, Address, City, Country, Code, Contact, Image}

func initializeRoutes() {

	router.Use(setUserStatus())

	//afrecast.html
	router.GET(routes[0], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"afrecast.html",
			gin.H{
				"title": "Dot",
			},
		)
	})
	//beneficiariesSE.html
	router.GET(routes[1], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			"beneficiariesSE.html",
			gin.H{
				"title": "Dot",
			},
		)
	})
	//cim.html
	router.GET(routes[2], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			"cim.html",
			gin.H{
				"title": "Dot",
			},
		)
	})
	//circle.html
	router.GET(routes[3], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			"circle.html",
			gin.H{
				"title": "Dot",
			},
		)
	})
	//dashboard.html
	router.GET(routes[4], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			"dashboard.html",
			gin.H{
				"title": "Dot",
			},
		)
	})
	//digitalbus.html
	router.GET(routes[5], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			"digitalbus.html",
			gin.H{
				"title": "Dot",
			},
		)
	})
	//dn.html
	router.GET(routes[6], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			"ewong.html",
			gin.H{
				"title": "Dot",
			},
		)
	})
	//ewong.html
	router.GET(routes[7], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			"ewong.html",
			gin.H{
				"title": "Dot",
			},
		)
	})
	//index.html
	router.GET(routes[8], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			"index.html",
			gin.H{
				"title": "Dot",
			},
		)
	})
	//jembe.html
	router.GET(routes[9], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			"jembe.html",
			gin.H{
				"title": "Dot",
			},
		)
	})
	//moow.html
	router.GET(routes[10], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			"moow.html",
			gin.H{
				"title": "Dot",
			},
		)
	})
	//notifications.html
	router.GET(routes[11], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			"nofitications.html",
			gin.H{
				"title": "Dot",
			},
		)
	})
	//programs.html
	router.GET(routes[12], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			"programs.html",
			gin.H{
				"title": "Dot",
			},
		)
	})
	//socialent.html
	router.GET(routes[13], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			"socialent.html",
			gin.H{
				"title": "Dot",
			},
		)
	})
	//survey.html
	router.GET(routes[14], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			"survey.html",
			gin.H{
				"title": "Dot",
			},
		)
	})
	//tables.html
	router.GET(routes[15], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			"tables.html",
			gin.H{
				"title": "Dot",
			},
		)
	})
	//tyds.html
	router.GET(routes[16], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			"tyds.html",
			gin.H{
				"title": "Dot",
			},
		)
	})
	//explore.html
	router.GET(routes[17], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			"explore.html",
			gin.H{
				"title": "Dot",
			},
		)
	})
	//user.html
	router.GET(routes[18], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			routes[17],

			gin.H{
				"title": "Dot",
				"user":  usermodel,
			},
		)
	})
	router.GET("/login", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			"sign-in.html",
			gin.H{
				"title": "Dot",
			},
		)
	})
	router.GET("register", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			"sign-up.html",
			gin.H{
				"title": "Dot",
			},
		)
	})

	userRoutes := router.Group("/u")
	{
		// Handle the GET requests at /u/login
		// Show the login page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/login", ensureNotLoggedIn(), showLoginPage)

		// Handle POST requests at /u/login
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/login", ensureNotLoggedIn(), performLogin)

		// Handle GET requests at /u/logout
		// Ensure that the user is logged in by using the middleware
		userRoutes.GET("/logout", ensureLoggedIn(), logout)

		// Handle the GET requests at /u/register
		// Show the registration page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/register", ensureNotLoggedIn(), showRegistrationPage)

		// Handle POST requests at /u/register
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/register", ensureNotLoggedIn(), register)
	}

	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			"sign-in.html",
			gin.H{
				"title": "Dot",
			},
		)
	})

}
