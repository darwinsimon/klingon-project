package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/darwinsimon/klingon-project/stapi"
	"github.com/darwinsimon/klingon-project/translator"
)

func init() {
	// Add timestamp and filename in the logging
	log.SetFlags(log.LstdFlags | log.Llongfile)

	flagV := flag.Bool("v", false, "print debug logs")
	flag.Parse()

	// Set global environment
	if *flagV {
		os.Setenv("PRINTLOG", "TRUE")
	}
}

func main() {

	// Get input parameter
	// Join multiple params as a single name
	var param = strings.Join(flag.Args(), " ")

	// Return explanation if there's no parameter
	if param == "" {
		fmt.Println("Please enter any name as parameter")
		fmt.Println("Example: ./klingon-project Uhura")
		return
	}

	// Translate name to klingon
	translated, err := translator.ToKlingon(param)
	if err != nil {
		// Failed
		fmt.Println("Your input name can't be translated")
	} else {
		// Successful
		fmt.Println()
		fmt.Println("Processing...")
		fmt.Println()

		// Print the species information
		speciesResult := getSpecies(param)

		fmt.Println("Klingon Name   :", translated)
		fmt.Println("Species        :", speciesResult)
		fmt.Println()

	}
}

func getSpecies(param string) string {
	stapiModule := stapi.Stapi{}
	character, charErr := stapiModule.CharacterSearch(param)

	// Successful -- print the species name
	if character != nil && charErr == stapi.ErrorNone {
		return character.Species
	}

	// Failed
	switch charErr {
	case stapi.ErrorTooManyResults:
		// Too many results
		return "Can't get species information. Your name is too common."
	case stapi.ErrorCharacterNotFound:
		// Character not found
		return "No information"
	default:
		return "An error occurred when retrieving the species information"
	}

}
