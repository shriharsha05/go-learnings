package main

import (
	"log"
	"os"
)

func main() {

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	os.Stdout.Write(data)
}
