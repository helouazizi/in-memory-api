package memory

import (
	"encoding/json"
	"in-memory-api/models"
	"log"
	"net/http"
	"sync"
)

type Memory_Store struct {
	Artists   []models.Artist
	Locations []models.Location
	Dates     []models.Date
	Relations []models.Relation
	Mutex     sync.Mutex
}

func NewMemoryStore() *Memory_Store {
	return &Memory_Store{}
}

// Fetches API data and stores it in memory
func (s *Memory_Store) LoadData() {
	urls := map[string]string{
		"artists":   "https://groupietrackers.herokuapp.com/api/artists",
		"locations": "https://groupietrackers.herokuapp.com/api/locations",
		"dates":     "https://groupietrackers.herokuapp.com/api/dates",
		"relation":  "https://groupietrackers.herokuapp.com/api/relation",
	}

	for key, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("Error fetching %s: %v", key, err)
		}
		defer resp.Body.Close()

		switch key {
		case "artists":
			json.NewDecoder(resp.Body).Decode(&s.Artists)
		case "locations":
			json.NewDecoder(resp.Body).Decode(&s.Locations)
		case "dates":
			json.NewDecoder(resp.Body).Decode(&s.Dates)
		case "relation":
			json.NewDecoder(resp.Body).Decode(&s.Relations)
		}
	}
	//fmt.Println(s.Artists)

	log.Println("Data loaded successfully!")
}
