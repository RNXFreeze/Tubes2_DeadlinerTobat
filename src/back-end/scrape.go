package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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

	// map[tier]map[element][][]string
	output := make(map[string]map[string][][]string)

	currentTier := ""

	doc.Find("h2, h3, tr").Each(func(i int, s *goquery.Selection) {
		if goquery.NodeName(s) == "h2" || goquery.NodeName(s) == "h3" {
			heading := strings.TrimSpace(s.Text())
			if strings.HasPrefix(heading, "Tier") {
				parts := strings.Fields(heading)
				if len(parts) >= 2 {
					currentTier = strings.ToLower(parts[0] + " " + parts[1])
					if _, exists := output[currentTier]; !exists {
						output[currentTier] = make(map[string][][]string)
					}
				}
			}
			return
		}

		if goquery.NodeName(s) != "tr" || currentTier == "" {
			return
		}

		left := s.Find("td").First()
		right := left.Next()

		element := strings.TrimSpace(left.Find("a").Text())
		if element == "" {
			return
		}

		var recipes [][]string
		right.Find("li").Each(func(j int, li *goquery.Selection) {
			var parts []string
			li.Find("a").Each(func(k int, a *goquery.Selection) {
				text := strings.TrimSpace(a.Text())
				if text != "" {
					parts = append(parts, text)
				}
			})
			if len(parts) == 2 {
				recipes = append(recipes, parts)
			}
		})

		if len(recipes) > 0 {
			output[currentTier][element] = recipes
		}
	})

	file, err := os.Create("output_by_tier.json")
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

	fmt.Println("Scraping completed. Data saved to output_by_tier.json")
}
