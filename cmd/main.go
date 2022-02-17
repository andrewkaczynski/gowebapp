package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bitly/go-simplejson"

	"github.com/gorilla/mux"
)

func getRoot(w http.ResponseWriter, r *http.Request) {

	hostname := os.Getenv("HOSTNAME")
	if hostname == "" {
		hostname = "not definied"
	}

	podIP := os.Getenv("POD_IP")
	if podIP == "" {
		podIP = "not definied"
	}

	containerImage := os.Getenv("CONTAINER_IMAGE")
	if containerImage == "" {
		containerImage = "not definied"
	}

	json := simplejson.New()
	json.Set("HOSTNAME", hostname)
	json.Set("POD_IP", podIP)
	json.Set("CONTAINER_IMAGE", containerImage)

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func main() {
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", getRoot)
	log.Fatal(http.ListenAndServe(":"+httpPort, router))
}
