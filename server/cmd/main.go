package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/devalparikh/EquityTrackr/server/pkg/datastore"

	"github.com/devalparikh/EquityTrackr/server/internal/investor"
	"github.com/devalparikh/EquityTrackr/server/internal/position"
)

type DBConnection = datastore.DBConnection

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequest(dbConnection DBConnection) {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/investors", investor.GetAllInvestors).Methods("GET")
	myRouter.HandleFunc("/investors/{name}", investor.GetInvestorById(dbConnection)).Methods("GET")
	myRouter.HandleFunc("/investors", investor.PostInvestor(dbConnection)).Methods("POST")

	myRouter.HandleFunc("/positions", position.PostPosition(dbConnection)).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(myRouter)

	log.Fatal(http.ListenAndServe(":8080", handler))
}

func main() {
	dbConnection := datastore.Run()
	datastore.Get(dbConnection, "investors")

	handleRequest(dbConnection)

	defer dbConnection.Client.Close()
}
