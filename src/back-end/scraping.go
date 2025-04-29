

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://little-alchemy.fandom.com/wiki/Elements_(Little_Alchemy_2)", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) Chrome/91.0.4472.124 Safari/537.36")
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	output := make(map[string][][]string)

	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		left := s.Find("td").First()
		right := left.Next()

		element := left.Find("a").Text()
		if element == "" {
			return 
		}

		var recipes [][]string

		right.Find("li").Each(func(j int, li *goquery.Selection) {
			parts := []string{}

			li.Find("a").Each(func(k int, a *goquery.Selection) {
				text := a.Text()
				if text != "" {
					parts = append(parts, text)
				}
			})

			if len(parts) == 2 {
				recipes = append(recipes, parts)
			}
		})

		if len(recipes) > 0 {
			output[element] = recipes
		}
	})

	file, err := os.Create("output.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(output)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Scraping completed. Data saved to output.json")
}
