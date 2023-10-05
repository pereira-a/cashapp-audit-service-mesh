package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Audit struct {
	Userid    string
	Operation string
	Metadata  string
	Timestamp time.Time
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/audit", GetAudit).
		Methods("GET").
		Queries("userid", "{userid}")

	r.HandleFunc("/audit", CreateAudit).
		Methods("POST")

	http.ListenAndServe(":8080", r)
}

func GetAudit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	audit := Audit{
		Userid:    "testUser",
		Operation: "sendMoney",
		Timestamp: time.Now(),
	}

	json.NewEncoder(w).Encode(audit)
}

func CreateAudit(w http.ResponseWriter, r *http.Request) {
	var user Audit
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Fprintf(w, "Received Audit: %s - %s - %s - %s", user.Userid, user.Operation, user.Metadata, user.Timestamp)
}
