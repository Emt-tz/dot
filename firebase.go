package main

import (
	"bytes"
	"context"
	"fmt"

	"io"
	"log"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"github.com/mitchellh/mapstructure"

	"google.golang.org/api/option"
)

//================================================platform user firebase function =============================================
func loaduser_by_email(id string) (string,string) {

	if err != nil {
		log.Fatalln("error initializing app: ", err)
	}
	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	load_user, _ := client.Collection("users").Doc(id).Get(ctx)

	mapstructure.Decode(load_user.Data(), &usermodel)
	return usermodel.Email,usermodel.Category

}

func loaduser_by_pass(id string) []byte {
	if err != nil {
		log.Fatalln("error initializing app: ", err)
	}
	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	load_user, _ := client.Collection("users").Doc(id).Get(ctx)
	mapstructure.Decode(load_user.Data(), &usermodel)

	return []byte(usermodel.Password)
}

func adduser(id string, data map[string]interface{}) error {
	
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

func edituser(id string, FirstName []string, LastName []string, Address []string, City []string, Country []string, Code []string) error {
	
	if err != nil {
		log.Fatalln("error initializing app: ", err)
	}
	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	_, err = client.Collection("users").Doc(id).Update(ctx, []firestore.Update{
		{Path: "FirstName", Value: FirstName},
		{Path: "LastName", Value: LastName},
		{Path: "Address", Value: Address},
		{Path: "City", Value: City},
		{Path: "Country", Value: Country},
		{Path: "Code", Value: Code},
	})

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
//================================================ end platform user firebase function =============================================

//================================================upload file firebase function ====================================================

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
//================================================upload file firebase function ====================================================


//===================================================programs firebase function ====================================================
//social innovation function is placed here
func get_social_beneficiaries_progress(Beneficiary string) map[string]interface{} {

	if err != nil {
		log.Fatalln("error initializing app: ", err)
	}
	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	iter := client.Collection("programs").Doc("SOCIAL-ENTREPRENEURSHIP").Collection(Beneficiary)

	res, err := iter.Doc("Progress").Get(ctx)
	if err != nil {
		fmt.Println(err)
	}

	return res.Data()

}
func update_social_beneficiaries_progress(Beneficiary string, Table string, data interface{}) error {
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

	iter := client.Collection("programs").Doc("SOCIAL-ENTREPRENEURSHIP").Collection(Beneficiary)

	_, err = iter.Doc("Progress").Update(ctx, []firestore.Update{
		{Path: Table, Value: data},
	})

	if err != nil {
		return err
	} else {
		return nil
	}

}

func social_beneficiaries_profile() map[string]interface{} {
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

	iter := client.Collection("programs").Doc("SOCIAL-ENTREPRENEURSHIP").Collection("Beneficiary")

	res, err := iter.Doc("Profile").Get(ctx)
	if err != nil {
		panic(err)
	}

	return res.Data()

}
//===================================================programs firebase function ====================================================
