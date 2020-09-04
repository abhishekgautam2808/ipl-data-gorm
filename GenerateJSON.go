package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	_ "database/sql"

	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
)

type TotalRunsOfTeams struct {
	Team string
	Runs int
}

type RunsByRcbPlayers struct {
	Batsman string
	Runs    int
}

type ForeignUmpireMatches struct {
	Country      string
	UmpiresCount int
}

type MatchesBySeason struct {
	Team    string
	Season  int
	Matches int
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "murdad@123"
	dbname   = "ipldb"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("connectiomn successful")

	var teamRuns []TotalRunsOfTeams
	err = db.Table("deliveries").Select("batting_team as team ,sum(total_runs) as runs").
		Group("batting_team").Find(&teamRuns).Error
	if err != nil {
		panic(err)
	}
	problem1(teamRuns)

	var playerRuns []RunsByRcbPlayers
	err = db.Table("deliveries").Select("batsman , sum(batsman_runs) as runs").Where("batting_team= ?", "Royal Challengers Bangalore").
		Group("batsman").Find(&playerRuns).Error
	if err != nil {
		panic(err)
	}
	problem2(playerRuns)

	var umpireMatches []ForeignUmpireMatches
	err = db.Table("umpires").Select("nationality as country, count(nationality) as umpires_count").
		Where("nationality <> ?", "India").Group("country").Find(&umpireMatches).Error
	if err != nil {
		panic(err)
	}
	problem3(umpireMatches)

	var teamMatches1 []MatchesBySeason

	err = db.Table("matches").Select("team1 as team, season , count(team1) as matches").
		Group("season,team1").Find(&teamMatches1).Error
	if err != nil {
		panic(err)
	}

	var teamMatches2 []MatchesBySeason
	err = db.Table("matches").Select("team2 as team, season , count(team2) as matches").
		Group("season,team2").Find(&teamMatches2).Error
	if err != nil {
		panic(err)
	}

	problem4(teamMatches1, teamMatches2)
}

func problem1(teamRuns []TotalRunsOfTeams) {
	dataMap1 := make(map[string]int)
	for _, row := range teamRuns {
		dataMap1[row.Team] = row.Runs
	}

	json1, err := json.Marshal(dataMap1)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("./web/output/json1.json", json1, 0644)
	if err != nil {
		panic(err)
	}
}

func problem2(playerRuns []RunsByRcbPlayers) {
	dataMap2 := make(map[string]int)
	for _, row := range playerRuns {
		dataMap2[row.Batsman] = row.Runs
	}

	json2, err := json.Marshal(dataMap2)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("./web/output/json2.json", json2, 0644)
	if err != nil {
		panic(err)
	}
}

func problem3(umpireMatches []ForeignUmpireMatches) {
	datamap3 := make(map[string]int)
	for _, row := range umpireMatches {
		datamap3[row.Country] = row.UmpiresCount
	}

	json3, err := json.Marshal(datamap3)
	if err != nil {
		panic(err)
	}
	_ = ioutil.WriteFile("./web/output/json3.json", json3, 0644)
}

func problem4(teamMatches1, teamMatches2 []MatchesBySeason) {

	datamap4 := make(map[string]map[int]int)
	for _, row := range teamMatches1 {
		v, ok := datamap4[row.Team]
		if ok {
			v[row.Season] = row.Matches
		} else {
			datamap4[row.Team] = map[int]int{row.Season: row.Matches}
		}
	}

	for _, row := range teamMatches2 {
		datamap4[row.Team][row.Season] += row.Matches

	}

	var years = []int{2008, 2009, 2010, 2011, 2012, 2013, 2014, 2015, 2016, 2017}

	for _, year := range years {
		for _, value := range datamap4 {
			_, ok := value[year]
			if ok {
				continue
			} else {
				value[year] = 0
			}
		}
	}

	json4, err := json.Marshal(datamap4)
	if err != nil {
		panic(err)
	}
	_ = ioutil.WriteFile("./web/output/json4.json", json4, 0644)
}
