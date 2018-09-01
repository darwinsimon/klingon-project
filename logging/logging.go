package logging

import (
	"log"
	"os"
)

// Println run log.Println() if PRINTLOG is TRUE
func Println(v ...interface{}) {
	if os.Getenv("PRINTLOG") == "TRUE" {
		log.Println(v)
	}
}
