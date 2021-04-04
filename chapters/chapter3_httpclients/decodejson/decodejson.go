package main

import (
	"encoding/json"
	"log"
	"net/http"
)

//decode json message after interacting with an endpoint

type Status struct {
	Message string
	Status  string
}

func main() {
	response, err := http.Post(
		"http://IP:PORT/ping",
		"application/json",
		nil,
	)
	if err != nil {
		log.Fatalln(err)
	}
	var status Status
	if err := json.NewDecoder(response.Body).Decode(&status); err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()
	log.Printf("%s -> %s\n", status.Status, status.Message)
}
