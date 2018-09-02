package resource

// Character of Star Trek
// Add showas to show the field's information to the console
type Character struct {
	UID         string `json:"uid"`
	Name        string `json:"name"`
	Species     string `json:"species" showas:"Species"`
	Gender      string `json:"gender"`
	YearOfBirth int    `json:"yearOfBirth"`
}
