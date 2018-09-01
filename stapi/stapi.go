package stapi

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/darwinsimon/klingon-project/logging"
)

var savedChar map[string]*Character

func init() {

	// Read from saved file
	if textData, err := ioutil.ReadFile("./char.txt"); err == nil {
		if err = json.Unmarshal(textData, &savedChar); err != nil {
			// Failed to read saved file -- initialize with empty savedChar
			savedChar = map[string]*Character{}
		}
	} else {
		// No saved file -- initialize with empty savedChar
		savedChar = map[string]*Character{}
	}
}

// CharacterSearch by name
func (s Stapi) CharacterSearch(name string) (*Character, StapiError) {

	// Path for character searching
	characterSearchPath := "/character/search"

	// Remove excess spaces and set it to lower case
	cleanName := strings.ToLower(strings.Trim(name, " "))

	// Set POST form for API call
	form := url.Values{}
	form.Add("name", cleanName)
	form.Add("title", cleanName)
	body := strings.NewReader(form.Encode())

	// Hit stapi.co API -- return ErrorCharacterNotFound if error
	response, err := s.restRequest(http.MethodPost, characterSearchPath, body)
	if err != nil {
		logging.Println("[ERR] CharacterSearch", err)
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
		logging.Println("[ERR] CharacterSearch", err)
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

		// Check the saved file before hitting stapi.co
		if savedChar[cleanName] != nil {
			logging.Println("[INFO] Use cached result for", name)
			return savedChar[cleanName], ErrorNone
		}

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

			// Assert information to selectedChar
			selectedChar.Species = detailChar.Species[0].Name
			selectedChar.UID = detailChar.UID

			// Update character saved file
			savedChar[cleanName] = &selectedChar
			encoded, err := json.Marshal(savedChar)
			if err != nil {
				// Encoding failed -- skip updating saved file
				logging.Println("[ERR] CharacterSearch", err)
			} else {
				if err = ioutil.WriteFile("char.txt", encoded, 0644); err != nil {
					// Saving file failed -- skip updating saved file
					logging.Println("[ERR] CharacterSearch", err)
				}
			}

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
		logging.Println("[ERR] getCharacter", err)
		return nil, ErrorCharacterNotFound
	}
	var result = struct {
		Character charResponse `json:"character"`
	}{}

	// Convert API response to struct -- return ErrorCharacterNotFound if error
	if err = json.Unmarshal(response, &result); err != nil {
		logging.Println("[ERR] getCharacter", err)
		return nil, ErrorCharacterNotFound
	}

	return &result.Character, ErrorNone
}

// Hit Stapi.co REST API
func (s Stapi) restRequest(method string, path string, body io.Reader) ([]byte, error) {

	logging.Println("[INFO] Start requesting to", path)

	// Prepare request
	req, err := http.NewRequest(method, RESTURL+path, body)
	if err != nil {
		logging.Println("[ERR] restRequest", err)
		return nil, err
	}

	// Hit the API
	resp, err := client.Do(req)
	if err != nil {
		logging.Println("[ERR] restRequest", err)
		return nil, err
	}

	// Close body to prevent memory leak
	if resp != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != 200 {
		err = errors.New("Status Not OK")
		logging.Println("[ERR] restRequest", err)
		return nil, err
	}

	// Process the response
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logging.Println("[ERR] restRequest", err)
		return nil, err
	}

	logging.Println("[INFO] Response from", path)
	logging.Println("----------------------------------")
	logging.Println(string(content))
	logging.Println("----------------------------------")

	return content, err
}
