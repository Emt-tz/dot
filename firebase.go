package main

import (
	"bytes"
	"context"
	"io"
	"log"

	"cloud.google.com/go/storage"
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

func Upload(id string, fileInput []byte, fileName string) error {

	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}

	object := bucket.Object(fileName)
	writer := object.NewWriter(ctx)

	//Set the attribute
	writer.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens": id}
	defer writer.Close()

	if _, err := io.Copy(writer, bytes.NewReader(fileInput)); err != nil {
		return err
	}

	if err := object.ACL().Set(context.Background(), storage.AllUsers, storage.RoleReader); err != nil {
		return err
	}

	return nil

}
