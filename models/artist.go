package models

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"` // api
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	//Locations    string   `json:"locations"`    // api
	//ConcertDates string   `json:"concertDates"` // api
	//Relations    string   `json:"relations"`    // api
}
