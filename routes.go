package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var routes = []string{"afrecast.html", "beneficiariesSE.html", "cim.html", "circle.html", "dashboard.html", "digitalbus.html", "dn.html", "ewong.html", "index.html", "jembe.html", "moow.html", "notifications.html", "programs.html", "socialent.html", "survey.html", "tables.html", "tyds.html", "user.html", "explore.html", "sign-in.html", "sign-up.html"}

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

var Title string
var Email string
var FirstName string
var LastName string
var Address string
var City string
var Country string
var Code string
var Contact string
var Image string

var usermodel user

func initializeRoutes() {

	router.GET(routes[0], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			routes[0],
			gin.H{
				"title": "Dot",
			},
		)
	})
	router.GET(routes[1], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			routes[1],
			gin.H{
				"title": "Dot",
			},
		)
	})
	router.GET(routes[2], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			routes[2],
			gin.H{
				"title": "Dot",
			},
		)
	})
	router.GET(routes[3], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			routes[3],
			gin.H{
				"title": "Dot",
			},
		)
	})
	router.GET(routes[4], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			routes[4],
			gin.H{
				"title": "Dot",
			},
		)
	})
	router.GET(routes[5], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			routes[5],
			gin.H{
				"title": "Dot",
			},
		)
	})
	router.GET(routes[6], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			routes[6],
			gin.H{
				"title": "Dot",
			},
		)
	})
	router.GET(routes[7], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			routes[7],
			gin.H{
				"title": "Dot",
			},
		)
	})
	router.GET(routes[8], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			routes[8],
			gin.H{
				"title": "Dot",
			},
		)
	})
	router.GET(routes[9], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			routes[9],
			gin.H{
				"title": "Dot",
			},
		)
	})
	router.GET(routes[10], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			routes[10],
			gin.H{
				"title": "Dot",
			},
		)
	})
	router.GET(routes[11], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			routes[11],
			gin.H{
				"title": "Dot",
			},
		)
	})
	router.GET(routes[12], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			routes[12],
			gin.H{
				"title": "Dot",
			},
		)
	})
	router.GET(routes[13], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			routes[13],
			gin.H{
				"title": "Dot",
			},
		)
	})
	router.GET(routes[14], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			routes[14],
			gin.H{
				"title": "Dot",
			},
		)
	})
	router.GET(routes[15], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			routes[15],
			gin.H{
				"title": "Dot",
			},
		)
	})
	router.GET(routes[16], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			routes[16],
			gin.H{
				"title": "Dot",
			},
		)
	})
	//load user template
	router.GET(routes[17], func(c *gin.Context) {
		Title = "Manager"
		Email = "peterkelvin16@gmail.com"
		FirstName = "Emmanuel"
		LastName = "Mtera"
		Address = "Mbezi Jogoo"
		City = "Dar es salaam"
		Country = "Tanzania"
		Code = "14128"
		Contact = "0744568803"
		Image = "../assets/img/profile-img01.jpg"

		usermodel = user{Title, Email, FirstName, LastName, Address, City, Country, Code, Contact, Image}
		c.HTML(
			http.StatusOK,

			routes[17],

			gin.H{
				"title": "Dot",
				"user":  usermodel,
			},
		)
	})
	router.GET(routes[18], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			routes[18],
			gin.H{
				"title": "Dot",
			},
		)
	})
	router.GET(routes[19], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			routes[19],
			gin.H{
				"title": "Dot",
			},
		)
	})
	router.GET(routes[20], func(c *gin.Context) {
		c.HTML(
			http.StatusOK,

			routes[20],
			gin.H{
				"title": "Dot",
			},
		)
	})

}
