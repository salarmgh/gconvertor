package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var conf ConfigMap = ConfigMap{}

func init() {
	var err error
	conf, err = ConfigLoad()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/scale", scaleHandler)

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	log.Println("Running on " + conf.Host + ":" + conf.Port)
	err := http.ListenAndServe(conf.Host + ":" + conf.Port, loggedRouter)
	if err != nil {
		log.Println(err)
	}
}
