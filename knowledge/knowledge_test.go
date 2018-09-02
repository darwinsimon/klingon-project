package knowledge_test

import (
	"fmt"
	"reflect"
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

func TestParseInfoValueString(t *testing.T) {

	value := reflect.ValueOf("String Value")
	assert.Equal(t, "String Value", ParseInfoValue(value))

}
func TestParseInfoValueEmptyString(t *testing.T) {

	value := reflect.ValueOf("")
	assert.Equal(t, "No information", ParseInfoValue(value))

}

func TestParseInfoValueIntZero(t *testing.T) {

	value := reflect.ValueOf(0)
	assert.Equal(t, "No information", ParseInfoValue(value))

}

func TestParseInfoValueInt(t *testing.T) {

	value := reflect.ValueOf(123)
	assert.Equal(t, "123", ParseInfoValue(value))

}

func TestParseInfoValueUnknownType(t *testing.T) {

	value := reflect.ValueOf(123.123)
	assert.Equal(t, "", ParseInfoValue(value))

}
