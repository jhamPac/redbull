package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

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
		switch tt {
		case html.ErrorToken:
			return
		case html.StartTagToken:
			t := tokenizer.Token()
			if t.Data == "span" {
				for _, a := range t.Attr {
					switch a.Val {
					case "country-name":
						text := tokenizer.Next()
						if text == html.TextToken {
							s := string(tokenizer.Text())
							fmt.Printf("%s\n", strings.ToLower(s))
						}
					case "dial-code":
						text := tokenizer.Next()
						if text == html.TextToken {
							fmt.Printf("%s\n", tokenizer.Text())
						}
					}
				}
			}
		}
	}
}
