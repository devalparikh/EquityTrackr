package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/devalparikh/EquityTrackr/server/internal/datastore"
	"github.com/devalparikh/EquityTrackr/server/internal/investor"
)

type DBConnection = datastore.DBConnection

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequest(dbConnection DBConnection) {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/investors", investor.GetAllInvestors).Methods("GET")
	myRouter.HandleFunc("/investors/{name}", investor.GetInvestorByName(dbConnection)).Methods("GET")
	myRouter.HandleFunc("/investors", investor.PostAllInvestors).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	dbConnection := datastore.Run()
	datastore.Get(dbConnection, "investors")

	handleRequest(dbConnection)

	defer dbConnection.Client.Close()
}
