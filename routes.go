package main

import (
	"github.com/gin-gonic/gin"
)

func initializeRoutes() {

	router.Use(setUserStatus())

	adminRoutes := router.Group("auth")
	{
		adminRoutes.GET("/admin", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title": "Administrator",
				"Image": "../assets/img/anime3.png",
			}, "admin.html")
		})
	}

	userRoutes := router.Group("")
	{
		userRoutes.GET("/login", ensureNotLoggedIn(), showLoginPage)
		userRoutes.POST("/login", ensureNotLoggedIn(), performLogin)
		userRoutes.GET("/logout", ensureLoggedIn(), logout)
		userRoutes.GET("/register", showRegistrationPage)
		userRoutes.POST("/p/register", ensureNotLoggedIn(), register)

		userRoutes.GET("/admin", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title": "Administrator",
				"Image": "../assets/img/anime3.png",
			}, "admin.html")
		})

		userRoutes.GET("/dashboard", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title": "Dashboard",
				"Image": "../assets/img/anime3.png",
			}, "index.html")
		})
		userRoutes.GET("/statistics", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title": "Statistics",
				"Image": "../assets/img/anime3.png",
			}, "dashboard.html")
		})
		userRoutes.GET("/survey", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title": "Survey",
				"Image": "../assets/img/anime3.png",
			}, "survey.html")
		})
		userRoutes.GET("/user", ensureLoggedIn(), func(c *gin.Context) {
			FirstName := ""
			LastName := ""
			Address := ""
			City := ""
			Country := ""
			Code := ""

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
				"title": "Programs",
				"Image": "../assets/img/anime3.png",
			}, "programs.html")
		})
		userRoutes.GET("/table", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title": "Tables",
				"Image": "../assets/img/anime3.png",
			}, "tables.html")
		})
		userRoutes.GET("/gallery", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title": "Gallery",
				"Image": "../assets/img/anime3.png",
			}, "gallery.html")
		})
		userRoutes.GET("/explore", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title": "Explore",
				"Image": "../assets/img/anime3.png",
			}, "explore.html")
		})
		//beneficiaries url
		userRoutes.GET("/socialbeneficiaries", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title": "Socialent",
				"Image": "../assets/img/anime3.png",
			}, "socialent.html")
		})
		userRoutes.GET("/digitaljob", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title": "Digital Job",
				"Image": "../assets/img/anime3.png",
			}, "digitaljob.html")
		})
		userRoutes.GET("/digitalbus", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title": "Digital Bus",
				"Image": "../assets/img/anime3.png",
			}, "digitalbus.html")
		})
		userRoutes.GET("/tyds", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title": "TYDS",
				"Image": "../assets/img/anime3.png",
			}, "tyds.html")
		})
		//beneficiaries routes
		userRoutes.GET("/beneficiariesSE.html", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title": "Beneficiaries",
				"Image": "../assets/img/anime3.png",
			}, "beneficiariesSE.html")
		})

		userRoutes.GET("/progress", SocialInnovators_profile_handler_get) //progress get
		userRoutes.GET("p/progress", SocialInnovators_progress_handler)   //progress post

	}
	router.GET("/", ensureNotLoggedIn(), func(c *gin.Context) {
		render(c, gin.H{
			"title": "Tables",
		}, "sign-in.html")
	})

	router.GET("/userdata", UserEdit)
	router.GET("/admino", func(c *gin.Context) {
		render(c, gin.H{
			"title": "Administrator",
			"Image": "../assets/img/anime3.png",
		}, "admin.html")
	})
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

	router.GET("/tre", func(c *gin.Context) {
		render(c, gin.H{
			"title": "Administrator",
			"Image": "../assets/img/anime3.png",
		}, "programstable.html")
	})

	router.GET("/test", func(c *gin.Context) {
		render(c, gin.H{
			"title": "Beneficiary Form",
			"Image": "../assets/img/anime3.png",
		}, "beneficiary_form.html")
	})

}
