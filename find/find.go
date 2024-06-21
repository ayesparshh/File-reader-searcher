package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type xkcd struct {
	Num        int    `json:"num"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
	Day 	  string `json:"day"`
	Month 	  string `json:"month"`
	Year 	  string `json:"year"`
}

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, " no file found ")
		os.Exit(-1)
	}

	fn := os.Args[1]
	
	var (
		input io.ReadCloser
		items []xkcd
		terms []string
		cnt   int
		err   error
	)

	input, err = os.Open(fn) 
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", fn, err)
		os.Exit(-1)
	} 

	if err = json.NewDecoder(input).Decode(&items) ; err != nil {
		fmt.Fprintf(os.Stderr, "Error reading JSON data: %v\n", err)
		os.Exit(-1)
	}

	fmt.Fprintf(os.Stderr, "Read %d comics\n", len(items))

	for _, t := range os.Args[2:] {
		terms = append(terms, strings.ToLower(t))
	}

outer:
	for _,item := range items {
		title := strings.ToLower(item.Title)
		transcript := strings.ToLower(item.Transcript)

		for _, term := range terms {
			if !strings.Contains(title, term) && !strings.Contains(transcript, term) {
				continue outer
			}
		}
		
		fmt.Printf("https://xkcd.com/%d/ %s/%s/%s  %q	\n", item.Num, item.Day, item.Month, item.Year, item.Title)
		cnt++	
	}


	if len(terms) < 1 {
		fmt.Fprintln(os.Stderr, "no search terms")
		os.Exit(-1)
	}

	fmt.Fprintf(os.Stderr, "found %d comics\n", cnt)

}	