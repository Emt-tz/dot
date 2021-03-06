// handlers.user.go

package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/mail"
	"strconv"
	// "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//==============================================================user login and sign up handlers =========================
func showLoginPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Login",
	}, "sign-in.html")
}

func user_category_urls(category string) string {
	if category == "Community Leader" {
		return "/digitalbusiness"
	} else if category == "Social Innovator" {
		return "/si"
	} else if category == "Digital Ambasador" {
		return "/digitaljob"
	} else {
		return "/tyds"
	}
}

func performLogin(c *gin.Context) {

	// Obtain the Posted username and password values
	email := c.PostForm("inputEmail")
	password := c.PostForm("inputPassword")

	category, err := isUserValid(email, password)
	// Check if the email/password combination is valid

	//now to differentiate admin from normal user

	if err == nil {
		if email == "admin@dot.com" {
			// If the username/password is valid set the token in a cookie
			token := generateSessionToken()
			c.SetCookie("token", token, 6000, "", "", false, true)
			c.SetCookie("admin", token, 6000, "","", false, true )
			c.Set("is_logged_in", true)
			c.Set("is_admin", true)
			// session.Save()
			c.Redirect(http.StatusMovedPermanently, "/admin")

		} else {
			path := user_category_urls(category)
			// If the username/password is valid set the token in a cookie
			token := generateSessionToken()
			c.SetCookie("token", token, 6000, "", "", false, true)
			c.Set("is_logged_in", true)
			c.Set("is_admin", false)
			// session.Save()
			c.Redirect(http.StatusMovedPermanently, path)
		}

	} else {
		// If the username/password combination is invalid,
		// show the error message on the login page
		c.HTML(http.StatusBadRequest, "sign-in.html", gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": err.Error() + "Invalid credentials provided"})
	}
}

func generateSessionToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}

func logout(c *gin.Context) {

	c.SetCookie("token", "", -1, "", "", true, true)
	c.SetCookie("admin", "", -1, "", "", true, true)

	render(c, gin.H{
		"title": "Login"}, "sign-in.html")

}

func showRegistrationPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Register"}, "sign-up.html")
}

func register(c *gin.Context) {
	// Obtain the POSTed username and password values
	username := c.PostForm("inputName")
	email := c.PostForm("inputEmail")
	password := c.PostForm("inputPassword")

	//parse email
	_, emailcheck := mail.ParseAddress(email)

	if emailcheck != nil {
		render(c, gin.H{
			"ErrorTitle":   "Failed",
			"ErrorMessage": emailcheck.Error()}, "sign-up.html")
	}

	if len(password) >= 8 {
		hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
		if err != nil {
			log.Fatalln(err)
		}
		Category := c.PostForm("Category")

		fmt.Println(hashedpassword)

		em, _ := loaduser_by_email(email)
		register_user := registerNewUser(email, username, email, hashedpassword, Category)

		if em != email {
			if register_user == nil {
				// token := generateSessionToken()
				// c.SetCookie("token", token, 6000, "", "", true, false)

				render(c, gin.H{
					"ErrorTitle":   "Registration Successfully",
					"ErrorMessage": "Log in to continue"}, "sign-in.html")

			} else {
				c.HTML(http.StatusBadRequest, "sign-up.html", gin.H{
					"ErrorTitle":   "Registration Failed",
					"ErrorMessage": "Try again"})
			}
		} else {
			c.HTML(http.StatusBadRequest, "sign-up.html", gin.H{
				"ErrorTitle":   "Registration Failed",
				"ErrorMessage": "Email Exists"})
		}

	} else {
		render(c, gin.H{
			"ErrorTitle":   "Short Password",
			"ErrorMessage": "Password must have 8 characters"}, "sign-up.html")
	}

}

func UserEdit(c *gin.Context) {
	response := make(map[string]int)
	// //get edited query parameters here
	queryparams := c.Request.URL.Query()

	Fname := queryparams["Fname"]
	Lname := queryparams["Lname"]
	Address := queryparams["Address"]
	City := queryparams["City"]
	Country := queryparams["Country"]
	Code := queryparams["Code"]

	if len(Fname) == 0 || len(Lname) == 0 || len(Address) == 0 || len(City) == 0 || len(Country) == 0 || len(Code) == 0 {
		response["uploaded"] = 0
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": response["uploaded"],
		})
	} else {
		v := edituser(usermodel.Email, Fname, Lname, Address, City, Country, Code)
		if v == nil {
			response["uploaded"] = 1
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": response["uploaded"],
			})
		} else {
			response["uploaded"] = 2
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": response["uploaded"],
			})
		}
	}

}

//=======================================================end user login and sign up handlers ==========================

//==============================================social innovation handlers =====================================================
//social innovators section handler
//social innovators table struct
type table struct {
	AIntervention  string
	BLead          string
	CDate          string
	DParticipation string
	EImpact        string
	FScoring       string
	GOutcome       string
}

//route is test
func SocialInnovators_progress_handler(c *gin.Context) {
	//this will handle both url queries
	if c.Query("edit") == "table1" && c.Query("action") == "edit" {
		if c.Query("Detail") != "" || c.Query("Value") != "" {

			tableid := c.Query("table") + c.Query("id")
			data := map[string]interface{}{
				c.Query("Detail"): c.Query("Value"),
			}
			response := update_social_beneficiaries_progress("Beneficiary", tableid, data)
			if response == nil {
				c.JSON(http.StatusOK, gin.H{"response": "submitted"})
			} else {
				c.JSON(http.StatusOK, gin.H{"response": response.Error()})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"response": "empty values to submit"})
		}
	} else if c.Query("edit") == "table" && c.Query("action") == "edit"{
		if c.Query("table") != "" {
			tableid := c.Query("id")

			data := table{
				AIntervention:  c.Query("Intervention"),
				BLead:          c.Query("Lead"),
				CDate:          c.Query("Date"),
				DParticipation: c.Query("Participation"),
				EImpact:        c.Query("Impact"),
				FScoring:       c.Query("Scoring"),
				GOutcome:       c.Query("Outcome"),
			}
			response := update_social_beneficiaries_progress("Beneficiary", c.Query("table")+tableid, data)

			if response == nil {
				c.JSON(http.StatusOK, gin.H{"response": "submitted"})
			} else {
				c.JSON(http.StatusOK, gin.H{"response": response.Error()})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"response": "empty values to submit"})
		}
	} else if c.Query("action") == "delete" {
		tableid := c.Query("id")
		response := delete_social_beneficiaries_progress("Beneficiary", c.Query("table")+tableid)
		if response == nil {
			c.JSON(http.StatusOK, gin.H{"response": "submitted"})
		} else {
			c.JSON(http.StatusOK, gin.H{"response": response.Error()})
		}
	}

}

//route is test
func SocialInnovators_progress_handler_get(c *gin.Context) {
	beneficiary := c.Query("beneficiary")
	response,err := get_social_beneficiaries_progress(beneficiary)
	if err == nil {
		render(c, gin.H{"title": "Profile"+beneficiary,
	"Image": "../assets/img/anime3.png","bene":response}, "siprofile.html")
	} else {
		c.JSON(http.StatusOK, gin.H{"response": err.Error()})
	}
	

}

func SocialInnovators_profile_handler(c *gin.Context) {
	beneficiary := c.Query("beneficiary")
	response,err := social_beneficiaries_profile(beneficiary)
	if err == nil {
		render(c, gin.H{
			"title":"Profile",
			"Image": "../assets/img/anime3.png",
			"data": response,
		}, "profile.html")
	}else{
		c.JSON(http.StatusOK, gin.H{"response": err.Error()})
	}
	// c.JSON(http.StatusOK, gin.H{"response": response})
}

//============================================== end social innovation handlers =====================================================
