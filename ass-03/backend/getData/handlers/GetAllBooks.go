package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"getData/models"
)

func (h handler) GetWeather(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var weathers []models.Weather

	if result := h.DB.Find(&weathers); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(weathers)
}
