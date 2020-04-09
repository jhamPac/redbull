package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/net/html"
)

func main() {
	data, err := ioutil.ReadFile("index.html")
	if err != nil {
		log.Fatalf("could not parse the file %v", err)
	}
	r := bytes.NewReader(data)
	tokenizer := html.NewTokenizer(r)

	for {
		tt := tokenizer.Next()
		if tt == html.ErrorToken {
			return
		}
		x := tokenizer.Token().Data
		fmt.Printf("token is: %v", x)
	}
}
