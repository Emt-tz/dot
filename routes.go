package main

import (
	"github.com/gin-gonic/gin"
)

func initializeRoutes() {

	router.Use(setUserStatus())

	userRoutes := router.Group("")
	{
		userRoutes.GET("/login", ensureNotLoggedIn(), showLoginPage)
		userRoutes.POST("/p/login", ensureNotLoggedIn(), performLogin)
		userRoutes.GET("/logout", ensureLoggedIn(), logout)
		userRoutes.GET("/register", showRegistrationPage)
		userRoutes.POST("/p/register", ensureNotLoggedIn(), register)

		userRoutes.GET("/dashboard", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title":    "Dashboard",
				"user_img": usermodel.Image,
			}, "index.html")
		})
		userRoutes.GET("/statistics", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title":    "Statistics",
				"user_img": usermodel.Image,
			}, "dashboard.html")
		})
		userRoutes.GET("/survey", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title":    "Survey",
				"user_img": usermodel.Image,
			}, "survey.html")
		})
		userRoutes.GET("/user", ensureLoggedIn(), func(c *gin.Context) {
			FirstName := usermodel.FirstName[0]
			if FirstName == "" {
				FirstName = ""
			} else {
				FirstName = usermodel.FirstName[0]
			}

			render(c, gin.H{
				"title":     usermodel.LastName,
				"id":        usermodel.Email,
				"Title":     usermodel.Title,
				"Email":     usermodel.Email,
				"FirstName": FirstName,
				"LastName":  usermodel.LastName,
				"Address":   usermodel.Address,
				"City":      usermodel.City,
				"Country":   usermodel.Country,
				"Code":      usermodel.Code,
				"Image":     "../assets/anime3.png",
				"Category":  usermodel.Category,
			}, "user.html")
		})
		userRoutes.GET("/programs", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title":    "Programs",
				"user_img": usermodel.Image,
			}, "programs.html")
		})
		userRoutes.GET("/table", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title":    "Tables",
				"user_img": usermodel.Image,
			}, "tables.html")
		})
		userRoutes.GET("/gallery", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title":    "Gallery",
				"user_img": usermodel.Image,
			}, "gallery.html")
		})
		userRoutes.GET("/explore", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title":    "Explore",
				"user_img": usermodel.Image,
			}, "explore.html")
		})

	}
	router.GET("/", ensureNotLoggedIn(), func(c *gin.Context) {
		render(c, gin.H{
			"title": "Tables",
		}, "sign-in.html")
	})

	router.GET("/userdata", UserEdit)
	router.GET("/session", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"title":     usermodel.LastName,
			"id":        usermodel.Email,
			"Title":     usermodel.Title,
			"Email":     usermodel.Email,
			"FirstName": usermodel.FirstName,
			"LastName":  usermodel.LastName,
			"Address":   usermodel.Address,
			"City":      usermodel.City,
			"Country":   usermodel.Country,
			"Code":      usermodel.Code,
			"Image":     "./assets/anime3.png",
			"Category":  usermodel.Category,
		})
	})

}
