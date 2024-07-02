package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Manuelshub/hng11_stage1/helpers"
	"net/http"
)

// Struct for the response in json
type Response struct {
	ClientIP string `json:"client_ip"`
	Location string `json:"location"`
	Greeting string `json:"greeting"`
}

type Failure struct {
	Message string `json:"message"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	visitorName := r.URL.Query().Get("client_name")
	if visitorName == "" {
		w.WriteHeader(http.StatusBadRequest)
		failure := Failure{Message: "visitor_name field is required"}
		json.NewEncoder(w).Encode(failure)
		return
	}

	clientIP := helpers.GetClientIp(r)
	clientCity := helpers.GetClientLocation(clientIP)
	temperature := helpers.GetTemperatureByCity(clientCity)
	greeting := fmt.Sprintf("Hello %s! The temperature is %f degrees Celcius in %s", visitorName, temperature, clientCity)

	response := Response{
		ClientIP: clientIP,
		Location: clientCity,
		Greeting: greeting,
	}
	json.NewEncoder(w).Encode(response)
}
