package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://animeflv.net/" // URL of the website

	// Make an HTTP GET request to the URL
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Parse the HTML using goquery
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the <strong> elements with class "Title"
	doc.Find("strong.Title").Each(func(index int, element *goquery.Selection) {
		animeTitle := element.Text()

		// Find the corresponding <span> element with class "Capi"
		episode := doc.Find("span.Capi").Eq(index).Text()

		fmt.Printf("Anime Title: %s Episode: %s\n", animeTitle, episode)
	})
}
