package models

import (
	"database/sql"
	"errors"
	"log"
	// "time"
)

type Link struct {
	Id       int     `json:"ispId"`
	SiteAX   float32 `json:"aSiteX"`
	SiteAY   float32 `json:"aSiteY"`
	SiteBX   float32 `json:"bSiteX"`
	SiteBY   float32 `json:"bSiteY"`
	FreqA    int     `json:"freqA"`
	FreqB    int     `json:"freqB"`
	Polar    string  `json:"polar"`
	Distance float32 `json:"distance"`
	Tech     string  `json:"technology"`
}

type LinkModel struct {
	DB *sql.DB
}

func (m *LinkModel) Get() (*[]Link, error) {
	query := `SELECT links.ISP_ID, links.frequency_A, links.frequency_B, links.polarization, links.distance, sites1.X_coordinate, sites1.Y_coordinate, sites2.X_coordinate, sites2.Y_coordinate, technologies.name
		  FROM links 
		  INNER JOIN sites AS sites1 ON links.site_A = sites1.ID 
		  INNER JOIN sites AS sites2 ON links.site_B = sites2.ID
		  INNER JOIN technologies ON links.technology = technologies.ID
		  WHERE links.is_active = 1`

	var linksSlice []Link
	links, err := m.DB.Query(query)
	for links.Next() {
		var l Link
		err = links.Scan(&l.Id, &l.FreqA, &l.FreqB, &l.Polar, &l.Distance, &l.SiteAX, &l.SiteAY, &l.SiteBX, &l.SiteBY, &l.Tech)
		if err != nil {
			log.Fatal(err)
		}
		linksSlice = append(linksSlice, l)
	}

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("models: no matching record found")
		} else {
			return nil, err
		}
	}

	return &linksSlice, nil
}
