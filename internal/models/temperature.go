package models

import (
	"database/sql"
	"errors"
	"log"
	// "time"
)

type TempImage struct {
	Time  string `json:"time"`
	Image string `json:"image"`
}

type TempImageModel struct {
	DB *sql.DB
}

func (m *TempImageModel) Get() (*[]TempImage, error) {
	query := `SELECT time, image_name FROM realtime_temperature_grids ORDER BY time DESC LIMIT 24`

	var tempImageSlice []TempImage
	temps, err := m.DB.Query(query)
	for temps.Next() {
		var t TempImage
		err = temps.Scan(&t.Time, &t.Image)
		if err != nil {
			log.Fatal(err)
		}
		tempImageSlice = append(tempImageSlice, t)
	}

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("models: no matching record found")
		} else {
			return nil, err
		}
	}

	return &tempImageSlice, nil
}
