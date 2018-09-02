package knowledge_test

import (
	"fmt"
	"testing"

	. "github.com/darwinsimon/klingon-project/knowledge"
	"github.com/stretchr/testify/assert"
)

func TestGetErrorTooManyResults(t *testing.T) {

	result := Get("ErrorTooManyResults", dummyRes)
	assert.Equal(t, "\nCan't get species information. Your name is too common.", result[0])
}
func TestGetErrorCharacterNotFound(t *testing.T) {

	result := Get("ErrorCharacterNotFound", dummyRes)
	assert.Equal(t, "\nThe system can't find any information regarding 'ErrorCharacterNotFound'", result[0])
}

func TestGetSuccessfulWithSpecies(t *testing.T) {

	result := Get(WithSpecies, dummyRes)
	assert.Equal(t, fmt.Sprintf("%-20s : %s", "Species", "Human"), result[0])
}

func TestGetSuccessfulWithBlankSpecies(t *testing.T) {

	result := Get("Foo", dummyRes)
	assert.Equal(t, fmt.Sprintf("%-20s : %s", "Species", "No information"), result[0])
}
