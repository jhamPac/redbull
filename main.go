package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

	// m := make(map[string]string)
	for {
		tt := tokenizer.Next()
		switch tt {
		case html.ErrorToken:
			fmt.Println("All done!")
			os.Exit(1)
		case html.StartTagToken:
			t := tokenizer.Token()
			if t.Data == "span" {
				for _, a := range t.Attr {
					switch a.Val {
					case "country-name":
						text := tokenizer.Next()
						if text == html.TextToken {
							s := string(tokenizer.Text())
							fmt.Printf("the country: %s\n", strings.ToLower(s))
							endtt := tokenizer.Next()
							if endtt == html.EndTagToken {
								n := tokenizer.Next()
								if n == html.StartTagToken {
									if tokenizer.Token().Data == "span" {
										text := tokenizer.Next()
										if text == html.TextToken {
											fmt.Printf("the next span text: %s\n", tokenizer.Text())
										}
									}

								}
							}
						}
					}
					fmt.Println("")
				}
			}
		}
	}
}
