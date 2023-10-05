package main

import (
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/audit", CreateAudit).
		Methods("GET").
		Queries("userid", "{userid}")

	// TODO: api post

	http.ListenAndServe(":8080", r)
}

func CreateAudit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["userid"]
	fmt.Fprintf(w, "You've requested a GET for userid=%s", user)
}
