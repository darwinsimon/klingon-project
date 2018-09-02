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

	// Set flags
	flagV := flag.Bool("v", false, "print debug logs")
	flag.Parse()

	// Set global environment enable verbose logging
	if *flagV {
		os.Setenv("PRINTLOG", "TRUE")
	}
}

func main() {
	fmt.Println()
	fmt.Println("KLINGON PROJECT")
	fmt.Println("---------------")
	fmt.Println()

	// Get input parameter
	// Join multiple params as a single name
	var param = strings.Join(flag.Args(), " ")

	// Return explanation if there's no parameter
	if param == "" {
		fmt.Println("Please enter any name as parameter")
		fmt.Println("Example: Uhura")
		fmt.Println()
		return
	}

	// Translate name to klingon
	translated, err := translator.ToKlingon(param)
	if err != nil {
		// Failed
		fmt.Println("Input          :", param)
		fmt.Println("Klingon Name   : Can't be translated")
		fmt.Println()
	} else {
		// Successful
		fmt.Println("Processing...")
		fmt.Println()

		// Print the species information
		speciesResult := getSpecies(param)

		fmt.Println("Input          :", param)
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
