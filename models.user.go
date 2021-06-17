// models.user.go

package main

import (
	"fmt"
	"log"
)

var email string
var username string
var Category string
var password string

// Check if the username and password combination is valid
func isUserValid(email, password string) error {
	user_login := loaduser_by_email(email)
	pass_word := loaduser_by_pass(email)
	if user_login == email && pass_word == password {
		fmt.Printf("Document data: %#v\n", user_login)
		return nil
	} else {
		return fmt.Errorf("something is not right")
	}

}

// Register a new user with the given username, email and password
func registerNewUser(id, username, email, password, Category string) error {
	data := map[string]interface{}{
		"id":        id,
		"Title":     username,
		"Email":     email,
		"FirstName": "",
		"Address":   "",
		"City":      "",
		"Country":   "",
		"Code":      "",
		"Contact":   "",
		"Image":     "",
		"Password":  password,
		"Category":  Category,
	}

	//add a user if and only if the user does not exist
	adduser := adduser(id, data)

	if adduser != nil {
		log.Fatal(adduser)
	}
	return nil

}
