// Package main is the entry point for this application.
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var name string

	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	log.Println(greet(name))
}

func greet(name string) string {
	if name == "" {
		return "Hello, World!"
	}

	return fmt.Sprintf("Hello, %s!", name)
}
