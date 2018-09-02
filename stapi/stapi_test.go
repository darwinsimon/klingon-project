package stapi_test

import (
	"os"
	"testing"

	. "github.com/darwinsimon/klingon-project/stapi"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

func TestCharacterSearchFailedToHitAPI(t *testing.T) {

	// Mock the API to reply 404
	defer gock.Off()
	gock.New(RESTURL).Post("/character/search").Reply(404)

	character, err := s.CharacterSearch("Foo")
	assert.Nil(t, character)
	assert.Equal(t, ErrorCharacterNotFound, err)

}

func TestCharacterSearchFailedToReadJSON(t *testing.T) {

	// Mock the API to reply wrong JSON
	defer gock.Off()
	gock.New(RESTURL).Post("/character/search").Reply(200).BodyString(`WRONG JSON RESPONSE`)

	character, err := s.CharacterSearch("Foo")
	assert.Nil(t, character)
	assert.Equal(t, ErrorCharacterNotFound, err)

}

func TestCharacterSearchExceedLimit(t *testing.T) {

	bodyString := `{"page":{"pageNumber":0,"pageSize":50,"numberOfElements":11,"totalElements":11,"totalPages":1,"firstPage":true,"lastPage":true},"sort":{"clauses":[]},"characters":[{"uid":"UID1","name":"Name 1"},{"uid":"UID2","name":"Name 2"},{"uid":"UID3","name":"Name 3"},{"uid":"UID4","name":"Name 4"},{"uid":"UID5","name":"Name 5"},{"uid":"UID6","name":"Name 6"},{"uid":"UID7","name":"Name 7"},{"uid":"UID8","name":"Name 8"},{"uid":"UID9","name":"Name 9"},{"uid":"UID10","name":"Name 10"},{"uid":"UID11","name":"Name 11"}]}`

	// Mock the API to reply 11 results; more than default tolerance limit
	defer gock.Off()
	gock.New(RESTURL).Post("/character/search").Reply(200).BodyString(bodyString)

	character, err := s.CharacterSearch("foo")
	assert.Nil(t, character)
	assert.Equal(t, ErrorTooManyResults, err)

}

func TestCharacterSearchNotFound(t *testing.T) {

	bodyString := `{"page":{"pageNumber":0,"pageSize":50,"numberOfElements":0,"totalElements":0,"totalPages":0,"firstPage":true,"lastPage":true},"sort":{"clauses":[]},"characters":[]}`

	// Mock the API to reply no results
	defer gock.Off()
	gock.New(RESTURL).Post("/character/search").Reply(200).BodyString(bodyString)

	character, err := s.CharacterSearch("foo")
	assert.Nil(t, character)
	assert.Equal(t, ErrorCharacterNotFound, err)

}

func TestCharacterSearchFailedGetSpecies(t *testing.T) {

	searchBodyString := `{"page":{"pageNumber":0,"pageSize":50,"numberOfElements":2,"totalElements":2,"totalPages":1,"firstPage":true,"lastPage":true},"sort":{"clauses":[]},"characters":[{"uid":"CHMA0000132571","name":"Jean-Luc Picard","gender":"M","yearOfBirth":2305,"monthOfBirth":7,"dayOfBirth":13,"placeOfBirth":"La Barre, France, Earth","yearOfDeath":null,"monthOfDeath":null,"dayOfDeath":null,"placeOfDeath":null,"height":null,"weight":null,"deceased":null,"bloodType":null,"maritalStatus":"SINGLE","serialNumber":"SP-937-215","hologramActivationDate":null,"hologramStatus":null,"hologramDateStatus":null,"hologram":false,"fictionalCharacter":false,"mirror":false,"alternateReality":false},{"uid":"CHMA0000112984","name":"Jean-Luc Picard","gender":null,"yearOfBirth":null,"monthOfBirth":null,"dayOfBirth":null,"placeOfBirth":null,"yearOfDeath":null,"monthOfDeath":null,"dayOfDeath":null,"placeOfDeath":null,"height":null,"weight":null,"deceased":null,"bloodType":null,"maritalStatus":null,"serialNumber":null,"hologramActivationDate":null,"hologramStatus":null,"hologramDateStatus":null,"hologram":false,"fictionalCharacter":false,"mirror":false,"alternateReality":false}]}`

	// Mock the API to reply 2 results & 2 wrong get characters
	defer gock.Off()
	gock.New(RESTURL).Post("character/search").Reply(200).BodyString(searchBodyString)
	gock.New(RESTURL).Get("character").MatchParam("uid", "CHMA0000132571").Reply(404)
	gock.New(RESTURL).Get("character").MatchParam("uid", "CHMA0000112984").Reply(200).BodyString("WRONG JSON RESPONSE")

	character, err := s.CharacterSearch("Jean-Luc Picard")
	assert.Nil(t, character)
	assert.Equal(t, ErrorCharacterNotFound, err)

}

func TestCharacterSearchFound2Results(t *testing.T) {

	searchBodyString := `{"page":{"pageNumber":0,"pageSize":50,"numberOfElements":2,"totalElements":2,"totalPages":1,"firstPage":true,"lastPage":true},"sort":{"clauses":[]},"characters":[{"uid":"CHMA0000132571","name":"Jean-Luc Picard","gender":"M","yearOfBirth":2305,"monthOfBirth":7,"dayOfBirth":13,"placeOfBirth":"La Barre, France, Earth","yearOfDeath":null,"monthOfDeath":null,"dayOfDeath":null,"placeOfDeath":null,"height":null,"weight":null,"deceased":null,"bloodType":null,"maritalStatus":"SINGLE","serialNumber":"SP-937-215","hologramActivationDate":null,"hologramStatus":null,"hologramDateStatus":null,"hologram":false,"fictionalCharacter":false,"mirror":false,"alternateReality":false},{"uid":"CHMA0000112984","name":"Jean-Luc Picard","gender":null,"yearOfBirth":null,"monthOfBirth":null,"dayOfBirth":null,"placeOfBirth":null,"yearOfDeath":null,"monthOfDeath":null,"dayOfDeath":null,"placeOfDeath":null,"height":null,"weight":null,"deceased":null,"bloodType":null,"maritalStatus":null,"serialNumber":null,"hologramActivationDate":null,"hologramStatus":null,"hologramDateStatus":null,"hologram":false,"fictionalCharacter":false,"mirror":false,"alternateReality":false}]}`

	charBodyString := `{"character":{"uid":"CHMA0000132571","name":"Jean-Luc Picard","gender":"M","yearOfBirth":2305,"monthOfBirth":7,"dayOfBirth":13,"placeOfBirth":"La Barre, France, Earth","yearOfDeath":null,"monthOfDeath":null,"dayOfDeath":null,"placeOfDeath":null,"height":null,"weight":null,"deceased":null,"bloodType":null,"maritalStatus":"SINGLE","serialNumber":"SP-937-215","hologramActivationDate":null,"hologramStatus":null,"hologramDateStatus":null,"hologram":false,"fictionalCharacter":false,"mirror":false,"alternateReality":false,"characterSpecies":[{"uid":"SPMA0000026314","name":"Human","numerator":1,"denominator":1}]}}`

	// Mock the API to reply 2 results
	defer gock.Off()
	gock.New(RESTURL).Post("character/search").Reply(200).BodyString(searchBodyString)
	gock.New(RESTURL).Get("character").MatchParam("uid", "CHMA0000132571").Reply(200).BodyString(charBodyString)

	character, err := s.CharacterSearch("Jean-Luc Picard")
	assert.NotNil(t, character)
	assert.Equal(t, "Jean-Luc Picard", character.Name)
	assert.Equal(t, "Human", character.Species)
	assert.Equal(t, ErrorNone, err)

	// Remove testing file
	os.Remove("char.txt")

}

func TestCharacterSearchFoundFromCache(t *testing.T) {

	searchBodyString := `{"page":{"pageNumber":0,"pageSize":50,"numberOfElements":1,"totalElements":1,"totalPages":1,"firstPage":true,"lastPage":true},"sort":{"clauses":[]},"characters":[{"uid":"CHMA0000000000","name":"Darwin Simon","gender":"M"}]}`

	// Mock the API to reply 2 results
	defer gock.Off()
	gock.New(RESTURL).Post("character/search").Reply(200).BodyString(searchBodyString)

	character, err := s.CharacterSearch("Darwin")
	assert.NotNil(t, character)
	assert.Equal(t, "Darwin", character.Name)
	assert.Equal(t, "Human", character.Species)
	assert.Equal(t, ErrorNone, err)

	// Remove testing file
	os.Remove("char.txt")

}
