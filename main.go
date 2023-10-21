package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/gocolly/colly"
)

// Output file type and fields for URL, Title and Text for each page
type JSONoutput struct {
	Url   string `json:"url"`
	Title string `json:"title"`
	Text  string `json:"Text"`
}

// Scrape the data from the page using url string
func Scrape(urlstr string) (JSONoutput, []byte, error) {

	var jsonout JSONoutput
	var htmlbody []byte

	jsonout.Url = urlstr

	// HTTP client creation using colly
	// Collector manages the network communication and responsible for the execution of the attached callbacks while a collector job is running.
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	// Check status
	c.OnResponse(func(r *colly.Response) {
		log.Println("response received", r.StatusCode)
	})
	c.OnError(func(r *colly.Response, err error) {
		log.Println("error:", r.StatusCode, err)
	})

	// Get HTML content
	c.OnResponse(func(r *colly.Response) {
		htmlbody = r.Body
	})

	// Get title
	c.OnHTML("title", func(e *colly.HTMLElement) {
		jsonout.Title = e.Text
	})

	// Get text
	c.OnHTML("body", func(e *colly.HTMLElement) {
		jsonout.Text = e.Text
	})

	err := c.Visit(urlstr)
	if err != nil {
		return jsonout, nil, err
	}

	return jsonout, htmlbody, nil

}

// Write results of scrape to html file and jsonlist file
func WriteResults(jo JSONoutput, hb []byte) error {
	// first part: save wikipedia page html to wikipages directory
	wikiDir := "./wikipages"

	// Check if the directory exists
	_, err := os.Stat(wikiDir)
	if os.IsNotExist(err) {
		// Directory does not exist, so create it
		err := os.MkdirAll(wikiDir, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
		}
	}
	// Create the filename
	fn := path.Base(jo.Url)
	fileName := fmt.Sprintf("./wikipages/%s.html", fn)

	// Write the HTML data to the file
	err = os.WriteFile(fileName, hb, 0644)
	if err != nil {
		return err
	}

	// second part: write extracted text for the item
	// Open output file
	outjl, err := os.OpenFile("./go_items.jl", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer outjl.Close()

	// Marshal struct to json
	itemjs, err := json.Marshal(jo)
	if err != nil {
		return err
	}

	// Write the json data to the file
	_, err = outjl.WriteString(string(itemjs) + "\n")
	if err != nil {
		return err
	}

	return nil

}

func main() {

	// list of Wikipedia URLs
	urls := []string{
		"https://en.wikipedia.org/wiki/Robotics",
		"https://en.wikipedia.org/wiki/Robot",
		"https://en.wikipedia.org/wiki/Reinforcement_learning",
		"https://en.wikipedia.org/wiki/Robot_Operating_System",
		"https://en.wikipedia.org/wiki/Intelligent_agent",
		"https://en.wikipedia.org/wiki/Software_agent",
		"https://en.wikipedia.org/wiki/Robotic_process_automation",
		"https://en.wikipedia.org/wiki/Chatbot",
		"https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
		"https://en.wikipedia.org/wiki/Android_(robot)",
	}

	for _, nextUrl := range urls {
		fmt.Printf("Now scraping %s\n", nextUrl)
		jsonout, htmlbody, err := Scrape(nextUrl)
		if err != nil {
			log.Fatalf("Error scraping URL %s: %v", nextUrl, err)
		}

		err = WriteResults(jsonout, htmlbody)
		if err != nil {
			log.Fatalf("Error writing results: %v", err)
		}

	}
}
