package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jhampac/redbull/parsers"
)

func main() {
	data, err := ioutil.ReadFile("index.html")
	if err != nil {
		log.Fatalf("could not parse the file %v", err)
	}

	m := parsers.CountryHTML(data)

	fmt.Printf("All done!\n+%v", m)
}
