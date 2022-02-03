package datastore

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type DBConnection struct {
	Client *firestore.Client
	Ctx    context.Context
}

func Run() DBConnection {
	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile("/Users/devalparikh/Documents/Development/Github/EquityTrackr/server/equitytrackr-firebase-adminsdk-m70x2-91950eaf42.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	dbConnection := DBConnection{Client: client, Ctx: ctx}
	return dbConnection
}

func Get(dbConnection DBConnection, collectionName string) {

	fmt.Printf("Fetching collection %v...\n", collectionName)

	iter := dbConnection.Client.Collection(collectionName).Documents(dbConnection.Ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}
}

func GetOne(dbConnection DBConnection, collectionName string, documentName string) (map[string]interface{}, error) {

	fmt.Printf("Fetching collection %v for documentName %v...\n", collectionName, documentName)

	dsnap, err := dbConnection.Client.Collection(collectionName).Doc(documentName).Get(dbConnection.Ctx)
	if err != nil {
		fmt.Printf("Error trying to find %v! \n", documentName)
		return nil, err
	}
	m := dsnap.Data()
	fmt.Printf("DocumentName data: %#v\n", m)
	fmt.Printf("Found %v! \n", documentName)

	return m, nil
}

func SetOne(dbConnection DBConnection, collectionName string, documentName string, document interface{}) (interface{}, error) {

	fmt.Printf("Adding new document to collection %v...\n", collectionName)

	_, err := dbConnection.Client.Collection(collectionName).Doc(documentName).Set(dbConnection.Ctx, document)
	if err != nil {
		fmt.Printf("Failed adding document. Error: %v\n", err)
	}

	return document, err
}

func AddOne(dbConnection DBConnection, collectionName string, document interface{}) (string, interface{}, error) {

	fmt.Printf("Adding new document to collection %v...\n", collectionName)

	DocumentRef, _, err := dbConnection.Client.Collection(collectionName).Add(dbConnection.Ctx, document)
	if err != nil {
		fmt.Printf("Failed adding document. Error: %v\n", err)
	}

	return DocumentRef.ID, document, err
}
