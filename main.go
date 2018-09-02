package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/darwinsimon/klingon-project/knowledge"
	"github.com/darwinsimon/klingon-project/resource/stapi"
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

		// Print informations -- use stapi resource
		stapiModule := stapi.Stapi{}
		infos := knowledge.Get(param, stapiModule)

		fmt.Printf("%-20s : %s\n", "Input", param)
		fmt.Printf("%-20s : %s\n", "Klingon Name", translated)

		for i := range infos {
			fmt.Println(infos[i])
		}

		fmt.Println()

	}
}
