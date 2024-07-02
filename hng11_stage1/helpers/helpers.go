package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

// GetClientIp gets the Ip address of the requester
func GetClientIp(r *http.Request) string {
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		ips := strings.Split(forwarded, ",")
		return strings.TrimSpace(ips[0])
	}

	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}

func GetClientLocation(clientIp string) string {
	// Struct that will help us retrieve the city from the response
	var result struct {
		City string `json:"city"`
	}
	// Make a request to ip-api.com to get the client location details
	resp, err := http.Get(fmt.Sprintf("http://ip-api.com/json/%s", clientIp))
	if err != nil {
		return err.Error()
	}
	// Ensuring proper closure of the response body
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err.Error()
	}

	return result.City
}

func GetTemperatureByCity(city string) float64 {
	var client http.Client

	resp, err := client.Get(fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", os.Getenv("WEATHER_API_KEY"), city))
	if err != nil {
		return 0
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var data map[string]interface{}

		err := json.NewDecoder(resp.Body).Decode(&data)
		log.Println(data)
		if err == nil {
			if curr, ok := data["current"].(map[string]interface{}); ok {
				if temp, ok := curr["temp_c"].(float64); ok {
					return temp
				}
			}
		}
	}

	return 0
}
