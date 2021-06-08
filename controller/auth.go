package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("id", 12090292)
	session.Set("email", "test@gmail.com")
	session.Save()
	c.HTML(
		http.StatusOK,
		"sign-in.html",
		gin.H{
			"title": "Dot",
		},
	)
}
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.HTML(
		http.StatusOK,
		"sign-up.html",
		gin.H{
			"title": "Dot",
		},
	)
}
