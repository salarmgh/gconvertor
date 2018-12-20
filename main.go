package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	c, err := ConfigLoad()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/scale", scaleHandler)

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	log.Println("Running on " + c.Host + ":" + c.Port)
	err = http.ListenAndServe(c.Host + ":" + c.Port, loggedRouter)
	if err != nil {
		log.Println(err)
	}
}
