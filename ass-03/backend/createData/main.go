package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "captainrr12"
	dbname   = "weather_db"
)

var (
	db  *sql.DB
	err error
)

type Weather struct {
	ID    int
	wind  int
	water int
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Succesfully connected to database")

	uptimeTicker := time.NewTicker(15 * time.Second)

	for {
		select {
		case <-uptimeTicker.C:
			CreateData()
			fmt.Println("Data updated")
		}
	}

}

func CreateData() {
	weather := Weather{}
	waterValue := rand.Intn(100)
	windValue := rand.Intn(100)

	sqlStatement := `
	INSERT INTO weathers (water,wind)
	VALUES ($1, $2)
	Returning *
	`
	err = db.QueryRow(sqlStatement, waterValue, windValue).
		Scan(&weather.ID, &weather.water, &weather.wind)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Weather Data : %+v\n", weather)
}
