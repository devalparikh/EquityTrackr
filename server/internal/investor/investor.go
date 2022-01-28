package investor

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/devalparikh/EquityTrackr/server/internal/datastore"
)

type DBConnection = datastore.DBConnection

type investor struct {
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

type investors []investor

func GetAllInvestors(w http.ResponseWriter, r *http.Request) {
	investors := investors{
		{Name: "Deval", Balance: 11229.90},
		{Name: "Bob", Balance: 12319.90},
	}
	json.NewEncoder(w).Encode(investors)
}

func GetInvestorByName(dbConnection DBConnection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		collectionName := "investors"
		investorName := vars["name"]

		investor, err := datastore.GetOne(dbConnection, collectionName, investorName)

		if err != nil {
			errorMessage := fmt.Sprintf("Error recieved while fetching from firebase: %v", err)
			http.Error(w, errorMessage, http.StatusInternalServerError)
		} else {
			json.NewEncoder(w).Encode(investor)
		}

	}
}

func PostAllInvestors(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test POST endpoint worked")
}
