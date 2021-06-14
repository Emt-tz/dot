package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/option"
)

var app *firebase.App
var ctx context.Context
var client *firestore.Client

func loaduser_by_email(id string) string {
	var c user
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

	load_user, err := client.Collection("users").Doc(id).Get(ctx)
	mapstructure.Decode(load_user.Data(), &c)

	return c.Email

}

func loaduser_by_pass(id string) interface{} {
	var c user
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

	load_user, err := client.Collection("users").Doc(id).Get(ctx)
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
