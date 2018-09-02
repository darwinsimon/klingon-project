package knowledge_test

import (
	"github.com/darwinsimon/klingon-project/resource"
)

type DummyResource struct{}

const WithSpecies = "WithSpecies"

// CharacterSearch by name
func (d DummyResource) CharacterSearch(name string) (*resource.Character, resource.Error) {
	if name == "ErrorCharacterNotFound" {
		return nil, resource.ErrorCharacterNotFound
	} else if name == "ErrorTooManyResults" {
		return nil, resource.ErrorTooManyResults
	}

	char := resource.Character{}
	if name == WithSpecies {
		char.Species = "Human"
	}

	return &char, resource.ErrorNone
}
