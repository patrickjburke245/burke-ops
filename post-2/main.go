package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	var city string
	var state string
	var country string

	fmt.Println("Hi! Do you need to know the weather? If so, what city do you live in?")
	fmt.Scan(&city)
	fmt.Println("What state do you live in? Please provide it in 2-character form.")
	fmt.Scan(&state)
	fmt.Println("What country do you live in? Please provide it in 2-character form.")
	fmt.Scan(&country)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")

	response, err := http.Get(fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s,%s,%s&limit=1&appid=%s", city, state, country, apiKey))
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	apiResponse, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var geocodingResponseObject [1]geocodingResponse
	json.Unmarshal(apiResponse, &geocodingResponseObject)

	response2, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/3.0/onecall?lat=%s&lon=%s&appid=%s&units=imperial", strconv.FormatFloat(geocodingResponseObject[0].Latitude, 'f', 2, 64), strconv.FormatFloat(geocodingResponseObject[0].Longitude, 'f', 2, 64), apiKey))

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	apiResponse2, err := io.ReadAll(response2.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject2 weatherResponse
	json.Unmarshal(apiResponse2, &responseObject2)

	fmt.Printf("The current temperature is %s degrees Fahrenheit.\n", strconv.FormatFloat(responseObject2.Current.Temperature, 'f', 2, 64))
}

type geocodingResponse struct {
	Zip       string  `json:"zip"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Country   string  `json:"country"`
}

type weatherResponse struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Timezone  string  `json:"timezone"`
	Current   weatherResponseDetails
}

type weatherResponseDetails struct {
	Temperature float64 `json:"temp"`
	Humidity    int     `json:"humidity"`
	WindSpeed   float64 `json:"wind_speed"`
}
