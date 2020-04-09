package caffeine

import (
	"bytes"
	"strings"

	"golang.org/x/net/html"
)

// ParseCountryHTML receives and []byte, parses and returns a map
func ParseCountryHTML(h []byte) map[string]string {
	m := make(map[string]string)
	r := bytes.NewReader(h)
	tokenizer := html.NewTokenizer(r)

	done := false
	for !done {
		tt := tokenizer.Next()
		switch tt {
		case html.ErrorToken:
			done = true
		case html.StartTagToken:
			t := tokenizer.Token()
			if t.Data == "span" {
				for _, a := range t.Attr {
					var key string
					var value string
					switch a.Val {
					case "country-name":
						tt := tokenizer.Next()
						if tt == html.TextToken {
							s := strings.Trim(strings.ToLower(string(tokenizer.Text())), "")
							rs := strings.ReplaceAll(s, " ", "-")
							key = rs
							endtt := tokenizer.Next()
							if endtt == html.EndTagToken {
								tt := tokenizer.Next()
								if tt == html.StartTagToken {
									t := tokenizer.Token()
									if t.Data == "span" {
										tt := tokenizer.Next()
										if tt == html.TextToken {
											code := strings.Trim(string(tokenizer.Text()), "")
											value = code
										}
									}

								}
							}
							m[key] = value
						}
					}
				}
			}
		}
	}
	return m
}
