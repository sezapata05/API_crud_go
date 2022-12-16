package main

import (
	"gorilla_mux/api"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	// json.NewEncoder(w).Encode("{\"message\" : \"Hello Word\"}")

}

func main() {

	r := mux.NewRouter()

	// Create the api object
	a := &api.API{}

	// register the routers
	a.RegisterRoutes(r)

	r.HandleFunc("/", HandleIndex).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:    ":8081",
		Handler: r,
	}

	log.Println("Listening ...")
	srv.ListenAndServe()
	// log.Fatal(http.ListenAndServe(":8000", r))
}
