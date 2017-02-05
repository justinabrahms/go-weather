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
	WeatherTime                WeatherTime   `json:"FCTTIME"`
	ProbabilityOfPercipitation string        `json:"pop"`
	Temperature                EngMet        `json:"temp"`
	Dewpoint                   EngMet        `json:"dewpoint"`
	Condition                  string        `json:"condition"`
        Sky                        string        `json:"sky"`
	SkyDescription             string        `json:"wx"`
	WindSpeed                  EngMet        `json:"wspd"`
	WindDirection              Winddirection `json:"wdir"`
	Humidity                   string        `json:"humidity"`
	Feelslike                  EngMet        `json:"Feelslike"`
	Snow                       EngMet        `json:"snow"`
	ForecastDescriptionNumbers string        `json:"fctcode"`
        Preasure                   EngMet        `json:"mslp"`
}

type Winddirection struct {
	Direction string `json:"dir"`
	Degrees   string `json:"degrees"`
}

type WeatherTime struct {
	Epoch   string `json:"Epoch"`
	Hour    string `json:"Hour"`
	Day     string `json:"mday"`
	Mon     string `json:"mon"`
	Year    string `json:"year"`
	Weekday string `json:"weekday_name"`
	// TODO(justinlilly): Need to do datetime using seconds since epoch
}

type EngMet struct {
	// TODO(justinlilly): It would be nice to have stronger types here.
	English string
	Metric  string
}

type WeatherResult struct {
	Time                time.Time
	Temperature         EngMet
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
