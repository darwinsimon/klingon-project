package knowledge

import (
	"fmt"
	"reflect"

	"github.com/darwinsimon/klingon-project/resource"
)

// Get any available informations from API provider
func Get(param string, src resource.Resource) []string {
	character, charErr := src.CharacterSearch(param)

	infoArr := []string{}

	// Successful -- return any available information
	if character != nil && charErr == resource.ErrorNone {

		to := reflect.TypeOf(*character)
		vo := reflect.ValueOf(*character)

		// Get all fields with 'showas'
		for i := 0; i < to.NumField(); i++ {
			showAs := to.Field(i).Tag.Get("showas")

			// Only show if showas isn't blank
			if showAs != "" {
				fieldValue := reflect.Indirect(vo).FieldByName(to.Field(i).Name)
				var infoValue string

				switch fieldValue.Interface().(type) {

				case string:
					// String information
					infoValue = fieldValue.String()

					// Blank value
					if infoValue == "" {
						infoValue = "No information"
					}

				case int:
					// Number information
					infoValue = fmt.Sprintf("%d", fieldValue.Int())

					// Blank value
					if infoValue == "0" {
						infoValue = "No information"
					}
				}

				infoArr = append(infoArr, fmt.Sprintf("%-20s : %s", showAs, infoValue))
			}
		}

	} else {

		// Failed -- return error message
		switch charErr {
		case resource.ErrorTooManyResults:
			// Too many results
			infoArr = append(infoArr, "\nCan't get species information. Your name is too common.")
		case resource.ErrorCharacterNotFound:
			// Character not found
			infoArr = append(infoArr, fmt.Sprintf("\nThe system can't find any information regarding '%s'", param))
		}
	}

	return infoArr
}
