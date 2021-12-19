package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type NbaTeam struct {
	FullName string `json:"fullName"`
}
type NbaGame struct {
	SeasonYear   string  `json:"seasonYear"`
	Leage        string  `json:"league"`
	StartTime    string  `json:"startTimeUTC"`
	EndTime      string  `json:"endTimeUTC"`
	VisitingTeam NbaTeam `json:"vTeam"`
	HomeTeam     NbaTeam `json:"hTeam"`
}
type ApiJson struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Results int       `json:"results"`
	Games   []NbaGame `json:"games"`
}
type DailyGamesResponse struct {
	Api ApiJson `json:"api"`
}

var gameNumber = 1

func printDailyNbaGames(gameDay time.Time) {
	url := "https://api-nba-v1.p.rapidapi.com/games/date/" + gameDay.Format("2006-01-02")
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("x-rapidapi-host", "api-nba-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", os.Getenv("RAPID_API_TOKEN"))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("error getting nba data")
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var dailyGamesResponse DailyGamesResponse
	if err := json.Unmarshal(body, &dailyGamesResponse); err != nil {
		log.Fatal(err)
	}
	for _, game := range dailyGamesResponse.Api.Games {
		gameDayStartTime := convertUTCtoCentral(game.StartTime)
		if gameDayStartTime.Format("2006-01-02") == time.Now().Format("2006-01-02") {
			fmt.Printf("Game %d\n", gameNumber)
			fmt.Printf("    Home team: %s\n", game.HomeTeam.FullName)
			fmt.Printf("    Away team: %s\n", game.VisitingTeam.FullName)
			fmt.Printf("    Start time: %s\n", convertUTCtoCentral(game.StartTime))
			gameNumber++
		}
	}
}

func convertUTCtoCentral(utcString string) time.Time {
	loc, _ := time.LoadLocation("America/Chicago")
	parsedTime, err := time.Parse("2006-01-02T15:04:05.999999999Z07:00", utcString)
	if err != nil {
		fmt.Println(err)
	}
	return parsedTime.In(loc)
}

func main() {
	currentDay := time.Now()
	// due to rest api returning dates by UTC time, have to query today and tomorrow
	printDailyNbaGames(currentDay)
	printDailyNbaGames(currentDay.AddDate(0, 0, 1))
}
