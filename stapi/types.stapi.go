package stapi

import (
	"net/http"
	"time"
)

type StapiError int

// Use global client to prevent increasing open files
var client = http.Client{
	// Set 10 seconds timeout
	Timeout: time.Duration(10 * time.Second),
}

// Stapi REST main module
type Stapi struct {
}

// Character of Star Trek
type Character struct {
	Name    string
	Species string
}

// Rest API Response struct
type charResponse struct {
	UID     string            `json:"uid"`
	Name    string            `json:"name"`
	Species []speciesResponse `json:"characterSpecies"`
}

// Species of Star Trek character
type speciesResponse struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
}
