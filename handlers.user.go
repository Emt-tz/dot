// handlers.user.go

package main

import (
	"encoding/json"
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

func JsonUpload(c *gin.Context) {
	var data map[string]interface{}

	jsonfile :=
		`{
			"Unique ID for Youth Leader": "TZDLYL000003",
			"Project": "DL",
			"YL name": "violeth lupenza",
			"Region": "Dar Es Salaam",
			"Moodle username": "lupenzavioleth",
			"Moodle email id": "lupenzavioleth94@gmail.com",
			"Gender": "Female",
			"Personal email": "lupenzavioleth94@gmail.com",
			"Primary Phone number": "756915619",
			"YL Contract start date": "01/09/2017",
			"YL Contract end date ": "30/06/2018",
			"Role in programming": "Only CL",
			"Participated in Launchlab": "No",
			"Participated in Peer Mentorship Program": "No",
			"Deployed in Community?": "Yes",
			"YLP Course Completion": "Yes",
			"YLP Course Completion - This was done as online course and hence no F2F update needed": "",
			"Year": "2017",
			"Note: For Tanzania, contract start date is same as date of deployment. Contract end date is same as end of deployment date. And this means these youth are to be categorized as Community Leader. They may have done some Social Innovator work as well.": "",
			"# of community beneficiaries who participate in DOT social innovation impactathons": ""
		}`

	json.Unmarshal([]byte(jsonfile), &data)

	// c.JSON(http.StatusOK, gin.H{
	// 	"code":    http.StatusOK,
	// 	"message": data, // cast it to string before showing
	// })
	//completed here
	id := "Tanzania Youth Leader"

	result := jsonupload(id, data)

	if result != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": result.Error(), // cast it to string before showing
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "uploaded", // cast it to string before showing
		})
	}
}

func JsonRead(c *gin.Context) {
	programname := "Tanzania Youth Leader"
	file := loadprograms(programname)
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": file, // cast it to string before showing
	})
}

//handler to handle to beneficiaries upload
func AddBeneficiaries(c *gin.Context) {

	render(c, gin.H{"title": "beneficiaries",
		"Image": "../assets/img/anime3.png"}, "addbene.html")
}
