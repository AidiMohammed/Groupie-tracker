package models

// struct of data for API - locations part
type Locations struct {
	id        int      `json:"id"`
	locations []string `json:"locations"`
	dates     string   `json:"dates"`
}

// struct of data for API - dates part
type Dates struct {
	id    int      `json:"id"`
	dates []string `json:"dates"`
}

// struct of data for API - relation part
type Relations struct {
	id             int                 `json:"id"`
	datesLocations map[string][]string `json:"datesLocations"`
}

// struct of data for API - artists part
type Artist struct {
	id           int      	`json:"id"`
	name         string   	`json:"name"`
	image        string   	`json:"image"`
	members      []string 	`json:"members"`
	concertDates Dates     	`json:"concertDates"`
	locations    Locations 	`json:"locations"`
	relations    Relations  `json:"relations"`
	creationDate int      	`json:"creationDate"`
	firstAlbum   string   	`json:"firstAlbum"`
}