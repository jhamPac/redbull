package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jhampac/redbull/caffeine"
)

func main() {
	data, err := ioutil.ReadFile("index.html")
	if err != nil {
		log.Fatalf("could not parse the file %v", err)
	}

	m := caffeine.ParseCountryHTML(data)

	fmt.Printf("All done!\n+%v", m)
}
