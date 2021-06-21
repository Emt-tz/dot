// middleware.auth.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// This middleware ensures that a request will be aborted with an error
// if the user is not logged in
func ensureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// If there's an error or if the token is empty
		// the user is not logged in
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if !loggedIn {

			//if token expired just redirect to login instead of 401
			if token, err := c.Cookie("token"); err != nil || token == "" {
				c.Redirect(http.StatusMovedPermanently, "/login")
			} else {
				c.Redirect(http.StatusMovedPermanently, "/Dashboard")
			}

		}
	}
}

func ensureToregister() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if loggedIn {
			if token, err := c.Cookie("token"); err != nil || token != "" {
				c.Redirect(http.StatusMovedPermanently, "/register")
			} else {
				c.Redirect(http.StatusMovedPermanently, "/Dashboard")
			}

		}
	}
}

// This middleware ensures that a request will be aborted with an error
// if the user is already logged in
func ensureNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// If there's no error or if the token is not empty
		// the user is already logged in
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if loggedIn {
			if token, err := c.Cookie("token"); err != nil || token != "" {
				c.Redirect(http.StatusMovedPermanently, "/logout")
			} else {
				c.Redirect(http.StatusMovedPermanently, "/Dashboard")
			}

		}
	}
}



// This middleware sets whether the user is logged in or not
func setUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := c.Cookie("token"); err == nil || token != "" {
			c.Set("is_logged_in", true)
		} else {
			c.Set("is_logged_in", false)
		}
	}
}
