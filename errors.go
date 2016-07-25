package helpers

import (
	"log"
	"os"
)

// Simple Panic on error
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// Prints to stderr
func errorLog(e ...interface{}) {
	log.New(os.Stderr, "", 0).Println(e)
}