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
		userRoutes.GET("/", ensureNotLoggedIn(), showLoginPage)
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
		//===============================user routes here ============================//
		//this route is for all admin only
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
		
		//===============================end user routes here ========================//
		//===============================desired program routes ======================//
		//======================social innovator routes ==========//
		userRoutes.GET("/si", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title": "Socialent",
				"Image": "../assets/img/anime3.png",
			}, "sipage.html")
		})
		userRoutes.GET("/si/profile", ensureLoggedIn(), SocialInnovators_profile_handler_get) //progress get
		userRoutes.GET("/si/beneficiaries", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title": "Gallery",
				"Image": "../assets/img/anime3.png",
			}, "si_beneficiaries.html")
		})
		userRoutes.GET("/si/gallery", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title": "Gallery",
				"Image": "../assets/img/anime3.png",
			}, "sigallery.html")
		})
		//sipage user routes
		userRoutes.GET("/si/user", ensureLoggedIn(), func(c *gin.Context) {
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
			}, "siuser.html")
		})
		
		//======================end social innovator routes =======//
		//======================digital bussines routes ==========//
		userRoutes.GET("/digitalbusiness", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title": "Digital Bus",
				"Image": "../assets/img/anime3.png",
			}, "clpage.html")
		})
		userRoutes.GET("/digitalbusiness/gallery", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title": "Gallery",
				"Image": "../assets/img/anime3.png",
			}, "cl_gallery.html")
		})
		userRoutes.GET("/digitalbusiness/profile", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title": "Gallery",
				"Image": "../assets/img/anime3.png",
			}, "cl_profile.html")
		})
		userRoutes.GET("/digitalbusiness/survey", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title": "Gallery",
				"Image": "../assets/img/anime3.png",
			}, "cl_survey.html")
		})
		userRoutes.GET("/digitalbusiness/registration", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title": "Gallery",
				"Image": "../assets/img/anime3.png",
			}, "cl_form.html")
		})
		//cluser page
		userRoutes.GET("/digitalbusiness/user", ensureLoggedIn(), func(c *gin.Context) {
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
			}, "cluser.html")
		})
		//======================end digital bussines routes ==========//
		//======================digital job  routes ===============//
		userRoutes.GET("/digitaljob", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title": "Digital Job",
				"Image": "../assets/img/anime3.png",
			}, "djpage.html")
		})
		userRoutes.GET("/digitaljob/participants", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title": "Digital Job",
				"Image": "../assets/img/anime3.png",
			}, "dj_participant.html")
		})
		userRoutes.GET("/digitaljob/survey", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title": "Digital Job",
				"Image": "../assets/img/anime3.png",
			}, "dj_survey.html")
		})
		userRoutes.GET("/digitaljob/ambassador", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title": "Digital Job",
				"Image": "../assets/img/anime3.png",
			}, "dj_ambassador_profile.html")
		})
		//digitaljob user
		userRoutes.GET("/digitaljob/user", ensureLoggedIn(), func(c *gin.Context) {
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
			}, "djuser.html")
		})
		//======================end digital job  routes ===============//
		//======================tyds routes ======================//
		userRoutes.GET("/tyds", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title": "TYDS",
				"Image": "../assets/img/anime3.png",
			}, "tydspage.html")
		})
		userRoutes.GET("/tyds/gallery", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title": "TYDS",
				"Image": "../assets/img/anime3.png",
			}, "tyds_gallery.html")
		})
		userRoutes.GET("/tyds/survey", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title": "TYDS",
				"Image": "../assets/img/anime3.png",
			}, "tyds_survey.html")
		})
		userRoutes.GET("/tyds/profile", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title": "TYDS",
				"Image": "../assets/img/anime3.png",
			}, "tyds_bene_profile.html")
		})
		
		
		//tyds user
		userRoutes.GET("/tyds/user", ensureLoggedIn(), func(c *gin.Context) {
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
			}, "tydsuser.html")
		})
		//====================== end tyds routes ======================//
		//=============================== end desired program routes ======================//
		//beneficiaries routes
		userRoutes.GET("/beneficiariesSE.html", ensureLoggedIn(), func(c *gin.Context) {
			render(c, gin.H{
				"title": "Beneficiaries",
				"Image": "../assets/img/anime3.png",
			}, "beneficiariesSE.html")
		})

		//shared routes are place here
		userRoutes.GET("p/progress", ensureLoggedIn(), SocialInnovators_progress_handler) //progress post

	}

	//post or api routes placed here free to access
	router.GET("/userdata", UserEdit)
	router.GET("/tre", func(c *gin.Context) {
		render(c, gin.H{
			"title": "Administrator",
			"Image": "../assets/img/anime3.png",
		}, "programstable.html")
	})
	//test functions are placed right here
	router.GET("/test", func(c *gin.Context) {
		render(c, gin.H{
			"title": "Beneficiary Form",
			"Image": "../assets/img/anime3.png",
		}, "beneficiary_form.html")
	})

}
