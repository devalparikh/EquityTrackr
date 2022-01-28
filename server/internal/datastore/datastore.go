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

type dbConnection struct {
	Client *firestore.Client
	Ctx    context.Context
}

func Run() dbConnection {
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
	dbConnection := dbConnection{Client: client, Ctx: ctx}
	return dbConnection
}

func Get(dbConnection dbConnection, collectionName string) {

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
