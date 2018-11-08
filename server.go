package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func newServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/getUser/{userID}", func(w http.ResponseWriter, r *http.Request) {
		userID := mux.Vars(r)["userID"]
		body := getUser(fmt.Sprintf("https://reqres.in/api/users/%s", userID))
		bodyToJSON(body)
		saveToFile(body)
		fmt.Fprintf(w, string(body))
		w.WriteHeader(http.StatusOK)
	}).Methods("GET", "POST")

	server := negroni.Classic()
	server.UseHandler(router)

	http.ListenAndServe(":3000", server)
}
