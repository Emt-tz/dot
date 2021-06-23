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

type customer struct {
	Name    string
	Email   string
	Address *address
}

type address struct {
	Street1 string
	Street2 string
	City    string
	State   string
	Zip     string `form:"label=Postal Code"`
}

//social innovators table struct is placed here
type table struct {
	Date          string
	Impact        string
	Intervention  string
	Lead          string
	Outcome       string
	Participation string
	Scoring       string
}

//==============================================================user login and sign up handlers =========================
func showLoginPage(c *gin.Context) {

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Login",
	}, "sign-in.html")
}

func performLogin(c *gin.Context) {

	// Obtain the POSTed username and password values
	email := c.PostForm("inputEmail")
	password := c.PostForm("inputPassword")

	err := isUserValid(email, password)
	// Check if the email/password combination is valid

	//now to differentiate admin from normal user

	if err == nil {
		if email == "admin@dot.com" {
			// If the username/password is valid set the token in a cookie
			token := generateSessionToken()
			c.SetCookie("token", token, 3600, "", "", false, true)
			c.Set("is_logged_in", true)
			// session.Save()
			c.Redirect(http.StatusMovedPermanently, "auth/admin")

		} else {
			// If the username/password is valid set the token in a cookie
			token := generateSessionToken()
			c.SetCookie("token", token, 3600, "", "", false, true)
			c.Set("is_logged_in", true)
			// session.Save()
			c.Redirect(http.StatusMovedPermanently, "/dashboard")
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

		em := loaduser_by_email(email)
		register_user := registerNewUser(email, username, email, hashedpassword, Category)

		if em != email {
			if register_user == nil {
				// token := generateSessionToken()
				// c.SetCookie("token", token, 3600, "", "", true, false)

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
			"message": response["uploaded"], // cast it to string before showing
		})
	} else {
		v := edituser(usermodel.Email, Fname, Lname, Address, City, Country, Code)
		if v == nil {
			response["uploaded"] = 1
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": response["uploaded"], // cast it to string before showing
			})
		} else {
			response["uploaded"] = 2
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": response["uploaded"], // cast it to string before showing
			})
		}
	}

}

//=======================================================end user login and sign up handlers ==========================

//==============================================social innovation handlers =====================================================
//social innovators section handler
//route is test
func SocialInnovators_progress_handler(c *gin.Context) {
	//get the query parameters values

	tableid := c.Query("id")

	data := table{
		Date:          c.Query("Date"),
		Impact:        c.Query("Impact"),
		Intervention:  c.Query("Intervention"),
		Lead:          c.Query("Lead"),
		Outcome:       c.Query("Outcome"),
		Participation: c.Query("Participation"),
		Scoring:       c.Query("Scoring"),
	}
	response := update_social_beneficiaries_progress("Beneficiary", tableid, data)

	c.JSON(http.StatusOK, gin.H{"response": response})

}

func SocialInnovators_profile_handler(c *gin.Context) {
	response := social_beneficiaries_profile()
	c.JSON(http.StatusOK, gin.H{"response": response})
}

//============================================== end social innovation handlers =====================================================
