package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var nflGameNumber = 1

type NflTeamStatus struct {
	Name  string `json:"name"`
	Score string `json:"score"`
}
type NflGame struct {
	Date          string        `json:"date"`
	Venue         string        `json:"venue"`
	GameName      string        `json:"name"`
	ShortName     string        `json:"shortName"`
	AwayTeamScore NflTeamStatus `json:"awayTeam"`
	HomeTeamScore NflTeamStatus `json:"homeTeam"`
}
type NflResponse struct {
	Data []NflGame `json:"data"`
}

func convertUTCtoCentralNfl(utcString string) time.Time {
	loc, _ := time.LoadLocation("America/Chicago")
	parsedTime, err := time.Parse("2006-01-02T15:05Z", utcString)
	if err != nil {
		fmt.Println(err)
	}
	return parsedTime.In(loc)
}

func PrintTodayNflGames() {
	url := "https://nfl-schedule.p.rapidapi.com/v1/schedules"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "nfl-schedule.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "bede94ef7fmsh17857bbdb6c1d78p15cbcdjsn5acaa724ce20")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var nflResponse NflResponse
	if err := json.Unmarshal(body, &nflResponse); err != nil {
		log.Fatal(err)
	}
    fmt.Println()
	fmt.Println(strings.Repeat("🏈", 8))
    fmt.Println()
	for _, game := range nflResponse.Data {
		gameDayStartTime := convertUTCtoCentralNfl(game.Date)
		if gameDayStartTime.Format("2006-01-02") == time.Now().Format("2006-01-02") {
			fmt.Printf("Nfl Game %d\n", nflGameNumber)
			fmt.Printf("    Home team: %s\n", game.HomeTeamScore.Name)
			fmt.Printf("    Away team: %s\n", game.AwayTeamScore.Name)
			fmt.Printf("    Time: %s\n", gameDayStartTime)
			nflGameNumber++
		}
	}
}
