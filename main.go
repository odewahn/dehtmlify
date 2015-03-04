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
    t, _ := s.Html()
	s.ReplaceWithHtml(t)
  })

  out, _ := doc.Html()
  fmt.Println(out)
}



func main() {
	ExampleScrape("ch01.html")
}