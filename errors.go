package helpers

import (
	"log"
	"os"
	"fmt"
)

// Simple Panic on error
func ErrorCheckAndPanic(e error) {
	if e != nil {
		panic(e)
	}
}

func ErrorCheckAndPrint(e error) {
	if e != nil {
		fmt.Println("ERROR: " + e.Error())
	}
}


// Prints to stderr
func ErrorLog(e ...interface{}) {
	log.New(os.Stderr, "", 0).Printf("%+v", e)
}
