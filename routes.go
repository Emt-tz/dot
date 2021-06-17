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
				"Image": "../assets/img/anime3.png",
			}, "index.html")
		})
		userRoutes.GET("/statistics", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title":    "Statistics",
				"Image": "../assets/img/anime3.png",
			}, "dashboard.html")
		})
		userRoutes.GET("/survey", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title":    "Survey",
				"Image": "../assets/img/anime3.png",
			}, "survey.html")
		})
		userRoutes.GET("/user", ensureLoggedIn(), func(c *gin.Context) {
			FirstName := usermodel.FirstName[0]
			LastName := usermodel.LastName[0]
			Address := usermodel.Address[0]
			City := usermodel.City[0]
			Country := usermodel.Country[0]
			Code := usermodel.Code[0]

			if FirstName == "" || LastName == "" || Address == "" || City == "" || Country == "" || Code == "" {
				FirstName = ""
				LastName = ""
				Address = ""
				City = ""
				Country = ""
				Code = ""
			} else {
				FirstName = usermodel.FirstName[0]
				LastName = usermodel.LastName[0]
				Address = usermodel.Address[0]
				City = usermodel.City[0]
				Country = usermodel.Country[0]
				Code = usermodel.Code[0]
			}

			render(c, gin.H{
				"title":     usermodel.LastName,
				"id":        usermodel.Email,
				"Title":     usermodel.Title,
				"Email":     usermodel.Email,
				"FirstName": FirstName,
				"LastName":  LastName,
				"Address":   Address,
				"City":      City,
				"Country":   Country,
				"Code":      Code,
				"Image":     "../assets/img/anime3.png",
				"Category":  usermodel.Category,
			}, "user.html")
		})
		userRoutes.GET("/programs", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title":    "Programs",
				"Image": "../assets/img/anime3.png",
			}, "programs.html")
		})
		userRoutes.GET("/table", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title":    "Tables",
				"Image": "../assets/img/anime3.png",
			}, "tables.html")
		})
		userRoutes.GET("/gallery", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title":    "Gallery",
				"Image": "../assets/img/anime3.png",
			}, "gallery.html")
		})
		userRoutes.GET("/explore", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title":    "Explore",
				"Image": "../assets/img/anime3.png",
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
			"Image":     "../assets/img/anime3.png",
			"Category":  usermodel.Category,
		})
	})

}
