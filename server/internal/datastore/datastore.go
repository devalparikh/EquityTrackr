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

	fmt.Printf("fetching collection %v...\n", collectionName)

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

func GetOne(dbConnection DBConnection, collectionName string, document string) (map[string]interface{}, error) {

	fmt.Printf("fetching collection %v for document %v...\n", collectionName, document)

	dsnap, err := dbConnection.Client.Collection(collectionName).Doc(document).Get(dbConnection.Ctx)
	if err != nil {
		fmt.Printf("error trying to find %v! \n", document)
		return nil, err
	}
	m := dsnap.Data()
	fmt.Printf("Document data: %#v\n", m)
	fmt.Printf("found %v! \n", document)

	return m, nil
}

// func Add(collection string, payload string) {
// 	_, _, err = client.Collection(collection).Add(ctx, map[string]interface{}{
// 		"first":  "Alan",
// 		"middle": "Mathison",
// 		"last":   "Turing",
// 		"born":   1912,
// 	})
// 	if err != nil {
// 		log.Fatalf("Failed adding aturing: %v", err)
// 	}
// }
