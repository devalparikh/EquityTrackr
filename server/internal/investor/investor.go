package investor

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type investor struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type invesetors []investor

func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	invesetors := invesetors{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}

	fmt.Println("Endpoint Hit: All invesetors endpoint")
	json.NewEncoder(w).Encode(invesetors)
}

func PostAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test POST endpoint worked")
}
