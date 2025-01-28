package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/joho/godotenv"
)

type Meteo struct {
	Time  string `json:"time"`
	Image string `json:"image"`
}

type MeteoModel struct {
	DB *sql.DB
}

func (m *MeteoModel) Get() (*[]Meteo, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("INFLUXDB_TOKEN")
	org := os.Getenv("INFLUX_ORG")
	url := "http://data.telcorain.cz"
	client := influxdb2.NewClient(url, token)
	fmt.Println(client.QueryAPI(org))

	queryAPI := client.QueryAPI(org)
	query := `from(bucket: "realtime_gauges")
            |> range(start: -10m)
            |> filter(fn: (r) => r._measurement == "reference_vut")`
	results, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	for results.Next() {
		fmt.Println(results.Record())
	}
	if err := results.Err(); err != nil {
		log.Fatal(err)
	}

	var tempImageSlice []Meteo
	return &tempImageSlice, nil
}
