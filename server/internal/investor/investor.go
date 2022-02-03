package investor

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/devalparikh/EquityTrackr/server/pkg/datastore"
)

type DBConnection = datastore.DBConnection

type investor struct {
	Username string  `json:"username"`
	Balance  float64 `json:"balance"`
	Email    string  `json:"email"`
}

type investorResponse struct {
	ID string `json:"ID"`
	investor
}

type investors []investor

func GetAllInvestors(w http.ResponseWriter, r *http.Request) {
	investors := investors{
		{Username: "Deval", Balance: 11229.90},
		{Username: "Bob", Balance: 12319.90},
	}
	json.NewEncoder(w).Encode(investors)
}

func GetInvestorById(dbConnection DBConnection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)

		collectionName := "investors"
		investorId := vars["name"]

		investor, err := datastore.GetOne(dbConnection, collectionName, investorId)

		if err != nil {
			errorMessage := fmt.Sprintf("%v not found", investorId)
			http.Error(w, errorMessage, http.StatusNotFound)
		} else {
			w.Header().Set("Content-Type", "text/json; charset=utf-8")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			json.NewEncoder(w).Encode(investor)
		}
	}
}

func PostInvestor(dbConnection DBConnection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		decoder := json.NewDecoder(r.Body)
		var newInvestor investor
		err := decoder.Decode(&newInvestor)
		if err != nil {
			errorMessage := fmt.Sprintf("Incorrect request format %v. Error: %v", newInvestor, err)
			http.Error(w, errorMessage, http.StatusBadRequest)
		}

		ID, _, err := datastore.AddOne(dbConnection, "investors", newInvestor)
		if err != nil {
			errorMessage := fmt.Sprintf("Could not save investor: %v", newInvestor)
			http.Error(w, errorMessage, http.StatusBadRequest)
		} else {

			createdInvestor := investorResponse{investor: newInvestor, ID: ID}

			w.Header().Set("Content-Type", "text/json; charset=utf-8")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			json.NewEncoder(w).Encode(createdInvestor)
		}
	}
}
