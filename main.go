package main

import (
  "fmt"
  "log"
  "os"
  "github.com/PuerkitoBio/goquery"
)


// Here's what I want to do:
//   $("section").each(function() { $(this).replaceWith(this.innerHTML) }`



func ExampleScrape(fn string) {
  f, err := os.Open(fn)
  if err != nil {
    log.Fatal(err)
  }
  defer f.Close()

  doc, err := goquery.NewDocumentFromReader(f) 
  if err != nil {
    log.Fatal(err)
  }

  doc.Find("section").Each(func(i int, s *goquery.Selection) {
	 html, _ := s.Html()
	 s.AfterHtml(html)
  })

  doc.Find("section").Each(func(i int, s *goquery.Selection) {
	 s.Remove()
  })
  
  fmt.Println(doc.Html())

}



func main() {
	ExampleScrape("ch01.html")
}