package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	_ "database/sql"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

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

	fmt.Println("Successfully connected!")

	db.DropTableIfExists(&Deliveries{}, &Matches{}, &Umpires{})
	db.CreateTable(&Deliveries{}, &Matches{}, &Umpires{})

	file, err := os.Open("deliveries.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	r := csv.NewReader(file)

	tx := db.Begin()
	counter := 0
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		matchid, _ := strconv.Atoi(row[0])
		over, _ := strconv.Atoi(row[4])
		ball, _ := strconv.Atoi(row[5])
		issuperover, _ := strconv.Atoi(row[9])
		wideruns, _ := strconv.Atoi(row[10])
		byeruns, _ := strconv.Atoi(row[11])
		legbyeruns, _ := strconv.Atoi(row[12])
		noballruns, _ := strconv.Atoi(row[13])
		penaltyruns, _ := strconv.Atoi(row[14])
		batsmanruns, _ := strconv.Atoi(row[15])
		extraruns, _ := strconv.Atoi(row[16])
		totalruns, _ := strconv.Atoi(row[17])

		record := Deliveries{
			MatchID:         matchid,
			Inning:          row[1],
			BattingTeam:     row[2],
			BowlingTeam:     row[3],
			Over:            over,
			Ball:            ball,
			Batsman:         row[6],
			NonStriker:      row[7],
			Bowler:          row[8],
			IsSuperOver:     issuperover,
			WideRuns:        wideruns,
			ByeRuns:         byeruns,
			LegbyeRuns:      legbyeruns,
			NoballRuns:      noballruns,
			PenaltyRuns:     penaltyruns,
			BatsmanRuns:     batsmanruns,
			ExtraRuns:       extraruns,
			TotalRuns:       totalruns,
			PlayerDismissed: row[18],
			DismissalKind:   row[19],
			Fielder:         row[20],
		}
		if counter == 0 {
			counter++

			continue
		} else {
			tx.Create(&record)
		}
		counter++
	}
	tx.Commit()

	file2, err := os.Open("matches.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer file2.Close()
	r2 := csv.NewReader(file2)

	mx := db.Begin()
	counter = 0
	for {
		v, err := r2.Read()
		if err == io.EOF {
			break
		}
		id, _ := strconv.Atoi(v[0])
		season, _ := strconv.Atoi(v[1])
		dlapplied, _ := strconv.Atoi(v[9])
		winbyruns, _ := strconv.Atoi(v[11])
		winbywickets, _ := strconv.Atoi(v[12])

		matchesRecord := Matches{
			ID:            id,
			Season:        season,
			City:          v[2],
			Date:          v[3],
			Team1:         v[4],
			Team2:         v[5],
			TossWinner:    v[6],
			TossDecision:  v[7],
			Result:        v[8],
			DlApplied:     dlapplied,
			Winner:        v[10],
			WinByRuns:     winbyruns,
			WinByWickets:  winbywickets,
			PlayerOfMatch: v[13],
			Venue:         v[14],
			Umpire1:       v[15],
			Umpire2:       v[16],
			Umpire3:       v[17],
		}
		if counter == 0 {
			counter++
			continue
		} else {
			mx.Create(&matchesRecord)
		}
		counter++
	}
	mx.Commit()

	file3, err := os.Open("umpires.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer file3.Close()
	r3 := csv.NewReader(file3)

	ux := db.Begin()
	counter = 0
	for {
		v, err := r3.Read()
		if err == io.EOF {
			break
		}
		firstassociated, _ := strconv.Atoi(v[2])
		lastassociated, _ := strconv.Atoi(v[3])
		noofmatches, _ := strconv.Atoi(v[4])
		umpiresRecord := Umpires{
			Umpire:          v[0],
			Nationality:     v[1],
			FirstAssociated: firstassociated,
			LastAssociated:  lastassociated,
			NoOfMatches:     noofmatches,
		}
		if counter == 0 {
			counter++
			continue
		} else {
			ux.Create(&umpiresRecord)
		}
		counter++
	}
	ux.Commit()

}

//Deliveries is a custom type for table-deliveries
type Deliveries struct {
	// ID              int    `gorm:"not null;PRIMARY_KEY;AUTO_INCREMENT"`
	MatchID         int    `gorm:"not null"`
	Inning          string `gorm:"type:varchar(100);not null"`
	BattingTeam     string `gorm:"type:varchar(100);not null"`
	BowlingTeam     string `gorm:"type:varchar(100);not null"`
	Over            int    `gorm:"not null"`
	Ball            int    `gorm:"not null"`
	Batsman         string `gorm:"type:varchar(100);not null"`
	NonStriker      string `gorm:"type:varchar(100);not null"`
	Bowler          string `gorm:"type:varchar(100);not null"`
	IsSuperOver     int    `gorm:"not null"`
	WideRuns        int    `gorm:"not null"`
	ByeRuns         int    `gorm:"not null"`
	LegbyeRuns      int    `gorm:"not null"`
	NoballRuns      int    `gorm:"not null"`
	PenaltyRuns     int    `gorm:"not null"`
	BatsmanRuns     int    `gorm:"not null"`
	ExtraRuns       int    `gorm:"not null"`
	TotalRuns       int    `gorm:"not null"`
	PlayerDismissed string `gorm:"type:varchar(100);"`
	DismissalKind   string `gorm:"type:varchar(100);"`
	Fielder         string `gorm:"type:varchar(100);"`
}

// TableName overrides the table name
func (d Deliveries) TableName() string {
	return "deliveries"
}

//Matches is a custom type for table-Matches
type Matches struct {
	ID            int    `gorm:"not null;PRIMARY_KEY"`
	Season        int    `gorm:"not null"`
	City          string `gorm:"type:varchar(100);not null"`
	Date          string `gorm:"type:varchar(100);not null"`
	Team1         string `gorm:"type:varchar(100);not null"`
	Team2         string `gorm:"type:varchar(100);not null"`
	TossWinner    string `gorm:"type:varchar(100);not null"`
	TossDecision  string `gorm:"type:varchar(100);not null"`
	Result        string `gorm:"type:varchar(100);not null"`
	DlApplied     int    `gorm:"not null"`
	Winner        string `gorm:"type:varchar(100);not null"`
	WinByRuns     int    `gorm:"not null"`
	WinByWickets  int    `gorm:"not null"`
	PlayerOfMatch string `gorm:"type:varchar(100);not null"`
	Venue         string `gorm:"type:varchar(100);not null"`
	Umpire1       string `gorm:"type:varchar(100);"`
	Umpire2       string `gorm:"type:varchar(100);"`
	Umpire3       string `gorm:"size:100;"`
}

// TableName overrides the table name
func (m Matches) TableName() string {
	return "matches"
}

//Umpires is a custom type for table-Umpires
type Umpires struct {
	// ID              int    `gorm:"not null;PRIMARY_KEY;AUTO_INCREMENT"`
	Umpire          string `gorm:"type:varchar(100);not null"`
	Nationality     string `gorm:"type:varchar(100);not null"`
	FirstAssociated int    `gorm:"not null"`
	LastAssociated  int    `gorm:"not null"`
	NoOfMatches     int    `gorm:"not null"`
}

// TableName overrides the table name
func (u Umpires) TableName() string {
	return "umpires"
}
