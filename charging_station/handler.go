package main

import (
	"log"
	"net/http"
)

var url string

func handlerequests() {

	//sm := http.NewServeMux()
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/stations", Getstations)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
