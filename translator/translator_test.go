package translator_test

import (
	"testing"

	. "github.com/darwinsimon/klingon-project/translator"
	"github.com/stretchr/testify/assert"
)

func TestToKlingonDictionary(t *testing.T) {

	cases := []string{
		"a", "b", "ch", "d", "e", "gh", "h", "i", "j", "l", "m", "n", "ng", "o", "p", "q", "Q", "r", "s", "t", "tlh", "u", "v", "w", "y", "'", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ",", ".", " ",
	}

	expectedTranslated := []string{
		"0xF8D0", "0xF8D1", "0xF8D2", "0xF8D3", "0xF8D4", "0xF8D5", "0xF8D6", "0xF8D7", "0xF8D8", "0xF8D9", "0xF8DA", "0xF8DB", "0xF8DC", "0xF8DD", "0xF8DE", "0xF8DF", "0xF8E0", "0xF8E1", "0xF8E2", "0xF8E3", "0xF8E4", "0xF8E5", "0xF8E6", "0xF8E7", "0xF8E8", "0xF8E9", "0xF8F0", "0xF8F1", "0xF8F2", "0xF8F3", "0xF8F4", "0xF8F5", "0xF8F6", "0xF8F7", "0xF8F8", "0xF8F9", "0xF8FD", "0xF8FE", "0x0020",
	}

	for c := range cases {
		translated, err := ToKlingon(cases[c])
		assert.Equal(t, expectedTranslated[c], translated)
		assert.Nil(t, err)
	}

}

func TestToKlingonInvalidCharacter(t *testing.T) {

	cases := []string{
		"c", "g", "k", "x", "z", "!", "@", "#", "$", "%", "^", "&", "*",
	}

	for c := range cases {
		translated, err := ToKlingon(cases[c])
		assert.Equal(t, "", translated)
		assert.EqualError(t, err, "Not translatable")
	}

}

func TestToKlingonSuccess(t *testing.T) {

	translated, err := ToKlingon("Test")
	assert.Equal(t, "0xF8E3 0xF8D4 0xF8E2 0xF8E3", translated)
	assert.Nil(t, err)

	translated, err = ToKlingon("Uhura")
	assert.Equal(t, "0xF8E5 0xF8D6 0xF8E5 0xF8E1 0xF8D0", translated)
	assert.Nil(t, err)

	translated, err = ToKlingon("qQqQ")
	assert.Equal(t, "0xF8DF 0xF8E0 0xF8DF 0xF8E0", translated)
	assert.Nil(t, err)

	translated, err = ToKlingon("Playing")
	assert.Equal(t, "0xF8DE 0xF8D9 0xF8D0 0xF8E8 0xF8D7 0xF8DC", translated)
	assert.Nil(t, err)

	translated, err = ToKlingon("T'Challa")
	assert.Equal(t, "0xF8E3 0xF8E9 0xF8D2 0xF8D0 0xF8D9 0xF8D9 0xF8D0", translated)
	assert.Nil(t, err)

	translated, err = ToKlingon("tlhchngngngngngnnngngngng")
	assert.Equal(t, "0xF8E4 0xF8D2 0xF8DC 0xF8DC 0xF8DC 0xF8DC 0xF8DC 0xF8DB 0xF8DB 0xF8DC 0xF8DC 0xF8DC 0xF8DC", translated)
	assert.Nil(t, err)
}

func TestToKlingonNotTranslatable(t *testing.T) {

	translated, err := ToKlingon("%")
	assert.Equal(t, "", translated)
	assert.EqualError(t, err, "Not translatable")

	translated, err = ToKlingon("Foo!!")
	assert.Equal(t, "", translated)
	assert.EqualError(t, err, "Not translatable")

	translated, err = ToKlingon("Candy")
	assert.Equal(t, "", translated)
	assert.EqualError(t, err, "Not translatable")

	translated, err = ToKlingon("Zet")
	assert.Equal(t, "", translated)
	assert.EqualError(t, err, "Not translatable")
}
