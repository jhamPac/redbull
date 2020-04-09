package main

import (
	"io/ioutil"
	"log"
)

func main() {
	_, err := ioutil.ReadFile("index.html")
	if err != nil {
		log.Fatalf("could not parse the file %v", err)
	}
}
