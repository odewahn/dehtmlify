package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"os/exec"
)

func renderStartTag(t html.Token) string {
	ret := "<" + t.Data
	for _, a := range t.Attr {
		ret += fmt.Sprintf(" %s='%s' ", a.Key, a.Val)
	}
	ret += ">"
	return ret
}

func renderEndTag(t html.Token) string {
	return fmt.Sprintf("</%s>", t.Data)
}

func parseHtml(r io.Reader) string {
	ret := ""
	d := html.NewTokenizer(r)
	for {
		// token type
		tokenType := d.Next()
		if tokenType == html.ErrorToken {
			return ret
		}
		token := d.Token()
		switch tokenType {
		case html.StartTagToken:
			if token.Data != "section" {
				ret += renderStartTag(token)
			}
		case html.TextToken: // text between start and end tag
			ret += token.Data
		case html.EndTagToken: // </tag>
			if token.Data != "section" {
				ret += renderEndTag(token)
			}
		case html.SelfClosingTagToken:
		}
	}
	return ret
}


func processFile(fn string) {
	
	// open file and pull out the sections
	f, err := os.Open(fn + ".html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	stripped := parseHtml(f)
	
	// Write the new, sectionless data to a temp file
	tmpFile, err := ioutil.TempFile(os.TempDir(), "tmp")
	tmpFile.Write([]byte(stripped))
	defer os.Remove(tmpFile.Name())
	
	// Run pandoc and save the results in the current working directory
	_, lookErr := exec.LookPath("pandoc")
    if lookErr != nil {
        log.Fatal(lookErr)
    }

	cmd := exec.Command("pandoc", "-f", "html", "-t", "markdown_github", tmpFile.Name(), "-o", fn + ".md")
	output, err := cmd.CombinedOutput()
	if err != nil {
	    log.Fatal(fmt.Sprint(err) + ": " + string(output))
	}  
}


// Here's the pandoc command we need to run 
//   pandoc -f html -t markdown_github ch01-no-sections.html > out.md


func main() {

	base_dirs, _ := ioutil.ReadDir(".")

	for _, f := range base_dirs {
		if !f.IsDir() {
			fn := strings.Split(f.Name(),".")
			if fn[len(fn)-1] == "html" {
				fmt.Printf("Processing %s.html\n", fn[0])	
				processFile(fn[0])			
			}
		}
	}


	

}
