package stapi

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// CharacterSearch by name
func (s Stapi) CharacterSearch(name string) (*Character, StapiError) {

	// Path for character searching
	characterSearchPath := "/character/search"

	// Remove excess spaces and set it to lower case
	trimmedName := strings.ToLower(strings.Trim(name, " "))

	// Set POST form for API call
	form := url.Values{}
	form.Add("name", trimmedName)
	form.Add("title", trimmedName)
	body := strings.NewReader(form.Encode())

	// Hit stapi.co API -- return ErrorCharacterNotFound if error
	response, err := s.restRequest(http.MethodPost, characterSearchPath, body)
	if err != nil {
		log.Println("[ERR] CharacterSearch", err)
		return nil, ErrorCharacterNotFound
	}

	var result = struct {
		Page struct {
			Total int `json:"totalElements"`
		}
		Characters []charResponse `json:"characters"`
	}{}

	// Convert API response to struct -- return ErrorCharacterNotFound if error
	if err = json.Unmarshal(response, &result); err != nil {
		log.Println("[ERR] CharacterSearch", err)
		return nil, ErrorCharacterNotFound
	}

	// If the results have exceed the limit, return error
	if result.Page.Total > maxToleranceResult {
		return nil, ErrorTooManyResults
	}

	/**
	API will return multiple results
	System will get species information starting from the first result
	If the character has multiple species, system will pick the first species
	*/
	for c := range result.Characters {
		// Assert the name of the character
		var selectedChar = Character{
			Name: result.Characters[c].Name,
		}

		// Get species information -- skip if error
		detailChar, err := s.getCharacter(result.Characters[c].UID)
		if err != ErrorNone {
			continue
		}

		// The selected characters has species information
		if len(detailChar.Species) > 0 {

			// Assert from the first species
			selectedChar.Species = detailChar.Species[0].Name

			return &selectedChar, ErrorNone
		}

	}

	return nil, ErrorCharacterNotFound
}

// Get detail information about specific character
func (s Stapi) getCharacter(uid string) (*charResponse, StapiError) {

	// Path for get character
	characterSearchPath := "/character?uid=" + uid

	// Hit stapi.co API -- return ErrorCharacterNotFound if error
	response, err := s.restRequest(http.MethodGet, characterSearchPath, nil)
	if err != nil {
		log.Println("[ERR] getCharacter", err)
		return nil, ErrorCharacterNotFound
	}
	var result = struct {
		Character charResponse `json:"character"`
	}{}

	// Convert API response to struct -- return ErrorCharacterNotFound if error
	if err = json.Unmarshal(response, &result); err != nil {
		log.Println("[ERR] getCharacter", err)
		return nil, ErrorCharacterNotFound
	}

	return &result.Character, ErrorNone
}

// Hit Stapi.co REST API
func (s Stapi) restRequest(method string, path string, body io.Reader) ([]byte, error) {

	log.Println("Start requesting to", path)

	// Prepare request
	req, err := http.NewRequest(method, RESTURL+path, body)
	if err != nil {
		log.Println("[ERR] restRequest", err)
		return nil, err
	}

	// Hit the API
	resp, err := client.Do(req)
	if err != nil {
		log.Println("[ERR] restRequest", err)
		return nil, err
	}

	// Close body to prevent memory leak
	if resp != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != 200 {
		err = errors.New("Status Not OK")
		log.Println("[ERR] restRequest", err)
		return nil, err
	}

	// Process the response
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[ERR] restRequest", err)
		return nil, err
	}

	log.Println("Result for", path)
	log.Println("----------------------------------")
	log.Println(string(content))
	log.Println("----------------------------------")

	return content, err
}
