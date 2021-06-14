package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/option"
)

func loaduser_by_email(id string) string {
	if err != nil {
		log.Fatalln("error initializing app: ", err)
	}
	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	load_user, _ := client.Collection("users").Doc(id).Get(ctx)
	mapstructure.Decode(load_user.Data(), &c)

	return c.Email

}

func loaduser_by_pass(id string) interface{} {
	if err != nil {
		log.Fatalln("error initializing app: ", err)
	}
	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	load_user, _ := client.Collection("users").Doc(id).Get(ctx)
	mapstructure.Decode(load_user.Data(), &c)

	return c.Password
}

func adduser(id string, data map[string]interface{}) error {
	ctx = context.Background()
	//init firebase
	opt := option.WithCredentialsFile("firebase.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalln("error initializing app: ", err)
	}
	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	_, err = client.Collection("users").Doc(id).Set(ctx, data)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
