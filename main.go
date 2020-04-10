package main

import (
	"encoding/json"
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
	j, err := json.MarshalIndent(m, "", "	")
	if err != nil {
		log.Fatalf("could not successfully encode to json: %v", err)
	}

	fmt.Printf("All done!\n%s", string(j))
}
