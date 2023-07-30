
package utils

import (
	"fmt"
	"log"
)

// LogError logs the error message and exits the program
func LogError(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

// HandleError checks if an error occurred and logs the error message
func HandleError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}

// PanicOnError checks if an error occurred and panics with the error message
func PanicOnError(err error) {
	if err != nil {
		panic(fmt.Sprintf("Fatal error: %v", err))
	}
}

