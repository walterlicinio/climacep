package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
)

type Response struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

var validZipcode = regexp.MustCompile(`^\d{8}$`)

func handleClimateRequest(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	if !validZipcode.MatchString(cep) {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	location, err := fetchLocation(cep)
	if location == ", " {
		log.Printf("can not find zipcode: %v", err)
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}

	log.Printf("Location found: %s", location)

	latitude, longitude, err := fetchCoordinates(location)
	if err != nil {
		log.Printf("Error fetching coordinates: %v", err)
		http.Error(w, "error fetching coordinates", http.StatusInternalServerError)
		return
	}

	tempC, err := fetchTemperature(latitude, longitude)
	if err != nil {
		log.Printf("Error fetching temperature: %v", err)
		http.Error(w, "error fetching temperature", http.StatusInternalServerError)
		return
	}

	response := Response{
		TempC: tempC,
		TempF: tempC*1.8 + 32,
		TempK: tempC + 273.15,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
