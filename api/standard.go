package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/ping", PingHandler).Methods("POST")

	http.ListenAndServe(":8080", router)

}

func PingHandler(w http.ResponseWriter, r *http.Request) {

	// ctx := r.Context()
	// ctx, cancel := context.WithTimeout(ctx, time.Duration(5*time.Second))
	// defer cancel()
	query := r.URL.Query().Get("query")
	fmt.Println(query)
	body := r.Body

	var test map[string]interface{}
	err := json.NewDecoder(body).Decode(&test)

	if err != nil {
		fmt.Println(err)

		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode(test)
}
