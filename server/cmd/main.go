package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/devalparikh/EquityTrackr/server/internal/datastore"
	"github.com/devalparikh/EquityTrackr/server/internal/investor"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/investors", investor.GetAllArticles).Methods("GET")
	myRouter.HandleFunc("/investors", investor.PostAllArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	dbConnection := datastore.Run()
	datastore.Get(dbConnection, "investors")

	handleRequest()

	defer dbConnection.Client.Close()
}
