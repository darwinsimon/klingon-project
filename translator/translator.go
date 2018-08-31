package translator

import (
	"errors"
	"strings"
)

var dictionary = map[string]string{
	"a":   "0xF8D0",
	"b":   "0xF8D1",
	"ch":  "0xF8D2",
	"d":   "0xF8D3",
	"e":   "0xF8D4",
	"gh":  "0xF8D5",
	"h":   "0xF8D6",
	"i":   "0xF8D7",
	"j":   "0xF8D8",
	"l":   "0xF8D9",
	"m":   "0xF8DA",
	"n":   "0xF8DB",
	"ng":  "0xF8DC",
	"o":   "0xF8DD",
	"p":   "0xF8DE",
	"q":   "0xF8DF",
	"Q":   "0xF8E0",
	"r":   "0xF8E1",
	"s":   "0xF8E2",
	"t":   "0xF8E3",
	"tlh": "0xF8E4",
	"u":   "0xF8E5",
	"v":   "0xF8E6",
	"w":   "0xF8E7",
	"y":   "0xF8E8",
	"'":   "0xF8E9",
	"0":   "0xF8F0",
	"1":   "0xF8F1",
	"2":   "0xF8F2",
	"3":   "0xF8F3",
	"4":   "0xF8F4",
	"5":   "0xF8F5",
	"6":   "0xF8F6",
	"7":   "0xF8F7",
	"8":   "0xF8F8",
	"9":   "0xF8F9",
	",":   "0xF8FD",
	".":   "0xF8FE",
	" ":   "0x0020",
}

// ToKlingon translate any value to klingon
// If the value isn't translatable, error will be given
func ToKlingon(value string) (string, error) {

	translated := []string{}

	// Total length of the string
	length := len(value)

	cursor := 0

	// Loop until all values are validated
	for cursor < length {

		// Search for 2 characters
		// expected : ch - gh - ng
		if cursor+1 < length {

			translatedChar, ok := dictionary[strings.ToLower(value[cursor:cursor+2])]

			// Matched!
			if ok {
				// Save translated characters to array
				translated = append(translated, translatedChar)

				// Add 2 to cursor
				cursor += 2

				// Skip the remaining process
				continue
			}
		}

		// Search for 3 characters
		// expected : tlh
		if cursor+2 < length {
			translatedChar, ok := dictionary[strings.ToLower(value[cursor:cursor+3])]

			// Matched!
			if ok {
				// Save translated characters to array
				translated = append(translated, translatedChar)

				// Add 3 to cursor
				cursor += 3

				// Skip the remaining process
				continue
			}
		}

		// Convert []byte to string
		currentChar := string(value[cursor])

		// Search the dictionary
		translatedChar, ok := dictionary[currentChar]

		// If failed, search for the lowercase
		if !ok {

			translatedChar, ok = dictionary[strings.ToLower(currentChar)]

			// The character isn't translatable -- return error
			if !ok {
				return "", errors.New("Not translatable")
			}
		}

		// Save translated characters to array
		translated = append(translated, translatedChar)

		// Add 1 to cursor
		cursor++

	}

	// Seperate each character with space
	return strings.Join(translated, " "), nil

}
