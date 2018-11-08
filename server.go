package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func newServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/getPerson", func(w http.ResponseWriter, r *http.Request) {
		body := get("https://reqres.in/api/users/1")
		bodyToJSON(body)
		saveToFile(body)
		w.WriteHeader(http.StatusOK)
	}).Methods("GET", "POST")

	server := negroni.Classic()
	server.UseHandler(router)

	http.ListenAndServe(":3000", server)
}
