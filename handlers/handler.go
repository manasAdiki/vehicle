package handlers

import (
	"golang-api/charging_station"
	"log"
	"net/http"
)

var url string

func Handlerequests() {

	//sm := http.NewServeMux()
	http.HandleFunc("/", charging_station.HomePage)
	http.HandleFunc("/stations", charging_station.Getstations)

	log.Fatal(http.ListenAndServe(":9090", nil))

}
