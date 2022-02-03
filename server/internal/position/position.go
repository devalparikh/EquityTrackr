package position

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/devalparikh/EquityTrackr/server/pkg/datastore"
)

type position struct {
	InvestorID     string  `json:"investorID"`
	Name           string  `json:"name"`
	Location       string  `json:"location"`
	MarketValue    float64 `json:"marketValue"`
	InvestedAmount float64 `json:"InvestedAmount"`
}

type positionResponse struct {
	ID string `json:"ID"`
	position
}

type DBConnection = datastore.DBConnection

// TOOD make generic Post function
func PostPosition(dbConnection DBConnection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Extract request body
		decoder := json.NewDecoder(r.Body)
		var newPosition position
		err := decoder.Decode(&newPosition)
		if err != nil {
			errorMessage := fmt.Sprintf("Incorrect request format %v. Error: %v", newPosition, err)
			http.Error(w, errorMessage, http.StatusBadRequest)
		}

		// Check if given investorId exists
		investor, err := datastore.GetOne(dbConnection, "investors", newPosition.InvestorID)
		if err != nil {
			errorMessage := fmt.Sprintf("Investor %v not found", newPosition.InvestorID)
			http.Error(w, errorMessage, http.StatusNotFound)
			return
		} else {
			fmt.Printf("Found investor %v of position %v", investor, newPosition)
		}

		// Save position to firebase
		ID, _, err := datastore.AddOne(dbConnection, "positions", newPosition)
		if err != nil {
			errorMessage := fmt.Sprintf("Could not save investor: %v", newPosition)
			http.Error(w, errorMessage, http.StatusBadRequest)
		} else {

			createdPosition := positionResponse{position: newPosition, ID: ID}

			w.Header().Set("Content-Type", "text/json; charset=utf-8")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			json.NewEncoder(w).Encode(createdPosition)
		}
	}
}
