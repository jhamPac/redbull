package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/jhampac/redbull/caffeine"
)

func main() {
	data, err := ioutil.ReadFile("index.html")
	if err != nil {
		log.Fatalf("could not parse the file %v", err)
	}

	m := caffeine.ParseCountryHTML(data)
	j, err := json.MarshalIndent(m, "", "	")
	if err != nil {
		log.Fatalf("could not successfully encode to json: %v", err)
	}

	ioutil.WriteFile("countries.json", j, os.ModePerm)
}
