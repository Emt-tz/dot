// models.user.go

package main

import (
	"errors"
	"strings"
)

type users struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

// opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
// 	app, err := firebase.NewApp(context.Background(), nil, opt)
// 	if err != nil {
// 		return nil, fmt.Errorf("error initializing app: %v", err)
// 	}
var userList = []users{
	{Username: "admin", Email: "admin@dot.com", Password: "admin"},
}

// Check if the username and password combination is valid
func isUserValid(email, password string) bool {
	for _, u := range userList {
		if u.Email == email && u.Password == password {
			return true
		}
	}
	return false
}

// Register a new user with the given username, email and password
// NOTE: For this demo, we
func registerNewUser(username, email, password string) (*users, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("The password can't be empty")
	} else if !isUsernameAvailable(username) {
		return nil, errors.New("The username isn't available")
	}

	u := users{Username: username, Email: email, Password: password}

	userList = append(userList, u)

	return &u, nil
}

// Check if the supplied username is available
func isUsernameAvailable(username string) bool {
	for _, u := range userList {
		if u.Username == username {
			return false
		}
	}
	return true
}
