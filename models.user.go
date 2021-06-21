// models.user.go

package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

var username string
var Category string

// Check if the username and password combination is valid
func isUserValid(email string, password string) error {
	//compare hash of passwords then proceed

	user_login := loaduser_by_email(email)
	pass_word := loaduser_by_pass(email)
	userpass := bcrypt.CompareHashAndPassword(pass_word, []byte(password))
	if user_login == email && userpass == nil {
		fmt.Printf("Document data: %#v\n", user_login)
		return nil
	} else {
		return fmt.Errorf(userpass.Error())
	}

}

// Register a new user with the given username, email and password
func registerNewUser(id string, username string, email string, password []byte, Category string) error {

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
		"Password":  string(password),
		"Category":  Category,
	}

	//add a user if and only if the user does not exist
	adduser := adduser(id, data)

	if adduser != nil {
		log.Fatal(adduser)
	}
	return nil

}
