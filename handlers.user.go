// handlers.user.go

package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
	if err == nil {
		// If the username/password is valid set the token in a cookie
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", true, false)
		c.Set("is_logged_in", true)
		c.Redirect(http.StatusMovedPermanently, "/dashboard")

	} else {
		// If the username/password combination is invalid,
		// show the error message on the login page
		c.HTML(http.StatusBadRequest, "sign-in.html", gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid credentials provided"})
	}
}

func generateSessionToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}

func logout(c *gin.Context) {

	c.SetCookie("token", "", -1, "", "", false, true)

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
	Category := c.PostForm("Category")

	em := loaduser_by_email(email)
	register_user := registerNewUser(email, username, email, password, Category)

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

}
