package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Main struct {
	Temp float64 `json:"temp"`
}

type List struct {
	Dt   int64 `json:"dt"`
	Main Main  `json:"main"`
}

type Data struct {
	Cod     string `json:"cod"`
	Message int    `json:"message"`
	Cnt     int    `json:"cnt"`
	List    []List `json:"list"`
}

const (
	BASE_URL    = "https://api.openweathermap.org/data/2.5"
	FORMAT_TIME = "Mon, 02 Jan 2006"
)

func fetchData() (Data, error) {
	var client = &http.Client{}
	var data Data

	apiKey := os.Getenv("API_KEY")

	endpoint := fmt.Sprintf("%s/forecast?q=Jakarta&appid=%s", BASE_URL, apiKey)

	request, errrorRequest := http.NewRequest("GET", endpoint, nil)

	if errrorRequest != nil {
		return Data{}, errrorRequest
	}

	response, errorResponse := client.Do(request)

	if errorResponse != nil {
		return Data{}, errorResponse
	}

	defer response.Body.Close()

	errorParseJson := json.NewDecoder(response.Body).Decode(&data)

	if errorParseJson != nil {
		return Data{}, errorParseJson
	}

	return data, nil
}

func printWeather(data Data) {
	fmt.Println("Weather Forecast: ")

	var prevDate string
	for _, item := range data.List {
		timeStamp := item.Dt
		date := time.Unix(timeStamp, 0).Format(FORMAT_TIME)

		// Check if the date has changed
		if date != prevDate {
			temperature := item.Main.Temp
			fmt.Printf("%s: %.2f °C\n", date, temperature)
			prevDate = date
		}
	}
}

func main() {
	errorEnv := godotenv.Load()
	if errorEnv != nil {
		log.Fatal("Error loading .env file")
	}

	data, errorFetchData := fetchData()
	if errorFetchData != nil {
		log.Panic("ERR: ", errorFetchData)
	}

	printWeather(data)
}
