package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"os"
)

func renderStartTag(t html.Token) string {
	ret := "<" + t.Data
	for _,a := range t.Attr {
		ret += fmt.Sprintf(" %s='%s' ", a.Key, a.Val)
	}
	ret += ">"
	return ret
}

func renderEndTag(t html.Token) string {
	return fmt.Sprintf("</%s>", t.Data)
}

func parseHtml(r io.Reader) {
	d := html.NewTokenizer(r)
	for {
		// token type
		tokenType := d.Next()
		if tokenType == html.ErrorToken {
			return
		}
		token := d.Token()
		switch tokenType {
		case html.StartTagToken:
			if token.Data != "section" {
				fmt.Print(renderStartTag(token))
			}
		case html.TextToken: // text between start and end tag
			fmt.Print(token.Data)
		case html.EndTagToken: // </tag>
			if token.Data != "section" {
				fmt.Print(renderEndTag(token))				
			}
		case html.SelfClosingTagToken:
			

		}
	}
}

func main() {
	f, err := os.Open("ch01.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	parseHtml(f)

}
