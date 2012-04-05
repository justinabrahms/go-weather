package weather

import (
	"log"
	"net/http"
	"strings"
	"encoding/json"
	"time"
)

const baseApi = "http://api.wunderground.com/api"

type WeatherFeed struct {
	Forecast []HourlyForecast `json:"hourly_forecast"`
}

type HourlyForecast struct {
	WeatherTime WeatherTime `json:"FCTTIME"`
	ProbabilityOfPercipitation string `json:"pop"`
	Temperature Temperature `json:"temp"`
}

type WeatherTime struct {
	Epoch string
	// TODO(justinlilly): Need to do datetime using seconds since epoch
}

type Temperature struct {
	// TODO(justinlilly): It would be nice to have stronger types here.
	English string
	Metric string
}

type WeatherResult struct {
	Time time.Time
	Temperature Temperature
	PercipitationChance uint8
}


type Credentials struct {
	Key string
	client *http.Client
}

func CreateClient(Key string) Credentials {
	return Credentials{Key, &http.Client{}}
}

func (c Credentials) getApiUrlForFeature(Feature string) (string) {
	s := []string{baseApi, c.Key, Feature, "q"}
	return strings.Join(s, "/")
}

func (c Credentials) Get10DayForecast(Location string) (WeatherFeed) {
	var weather *WeatherFeed;
	url := strings.Join([]string{c.getApiUrlForFeature("hourly"), Location + ".json"}, "/")
	resp, err := c.client.Get(url)
	if err != nil {
		log.Println("Error fetching weather: %s for url: %s", err.Error(), url)
		return *weather
	}
	decoder := json.NewDecoder(resp.Body);
	if err = decoder.Decode(&weather); err != nil {
		log.Println("Error decoding weather feed: %s", err.Error())
		return *weather
	}
	return *weather
}
