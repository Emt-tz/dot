// models.user.go

package main

import (
	"fmt"
	"log"
)

//user logged in model
type user struct {
	Title     string
	Email     string
	Password  string
	FirstName string
	LastName  string
	Address   string
	City      string
	Country   string
	Code      string
	Contact   string
	Image     string
	Category  string
}

var Title = "Manager"
var Email = "peterkelvin16@gmail.com"
var FirstName = "Emmanuel"
var LastName = "Mtera"
var Address = "Mbezi Jogoo"
var City = "Dar es salaam"
var Country = "Tanzania"
var Code = "14128"
var Contact = "0744568803"
var Image = "../assets/img/profile-img01.jpg"

var usermodel = user{Title, Email, FirstName, "", LastName, Address, City, Country, Code, Contact, Image, "nil"}

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
