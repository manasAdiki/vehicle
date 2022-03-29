package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Getstations(w http.ResponseWriter, r *http.Request) {

	fuel_type := r.URL.Query().Get("fuel_type")
	state := r.URL.Query().Get("state")
	limit := r.URL.Query().Get("limit")

	url := "https://developer.nrel.gov/api/alt-fuel-stations/v1.json?api_key=DEMO_KEY&" + "fuel_type=" + fuel_type + "&state=" + state + "&limit=" + string(limit)
	fmt.Println(url)
	req, err := http.NewRequest(
		http.MethodGet, url, nil,
	)
	if err != nil {
		log.Fatalf("error creating HTTP request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error sending HTTP request: %v", err)
	}
	responseBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("error reading HTTP response body: %v", err)
	}

	errr := json.Unmarshal(responseBytes, &API)
	if errr != nil {
		log.Fatalf("error deserializing weather data")
	}

	json.NewEncoder(w).Encode(API)
	log.Printf("User Accesed records")

	fmt.Println("The fuel type ", fuel_type)
	fmt.Println("The state ", state)
	fmt.Println("Limit ", limit)
	// printing the access codes
	fmt.Println("The access codes are")
	for _, rec := range API.FuelStations {
		fmt.Println(rec.AccessCode)
	}
}
