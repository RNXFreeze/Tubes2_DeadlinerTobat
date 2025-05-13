/* Kelompok   : Kelompok 21 - Deadliner Tobat                                    */
/* Nama - 1   : Muhammad Raihan Nazhim Oktana                                    */
/* NIM - 1    : K01 - 13523021 - Teknik Informatika (IF-Ganesha) ITB             */
/* Nama - 2   : Mayla Yaffa Ludmilla                                             */
/* NIM - 2    : K01 - 13523050 - Teknik Informatika (IF-Ganesha) ITB             */
/* Nama - 3   : Anella Utari Gunadi                                              */
/* NIM - 3    : K02 - 13523078 - Teknik Informatika (IF-Ganesha) ITB             */
/* Tanggal    : Senin, 12 Mei 2025                                               */
/* Tugas      : Tugas Besar 2 - Strategi Algoritma (IF2211-24)                   */
/* File Path  : Tubes2_DeadlinerTobat/src/data/scrape.go                         */
/* Deskripsi  : F06 - Scraping (Website Little Alchemy 2)                        */
/* PIC F06    : K01 - 13523050 - Mayla Yaffa Ludmilla                            */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getMythElements() []string {
	var myth []string
	seen := make(map[string]struct{})

	url := "https://little-alchemy.fandom.com/wiki/Elements_(Myths_and_Monsters)"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	res, err := client.Do(req)
	if err != nil {
		log.Println("Myth HTTP error:", err)
		return myth
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Println("Myth status error:", res.StatusCode, res.Status)
		return myth
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println("Myth parse error:", err)
		return myth
	}

	doc.Find("div.mw-parser-output table tr").Each(func(_ int, row *goquery.Selection) {
		tdName := row.Find("td").Eq(0)
		name := strings.TrimSpace(tdName.Text())
		if name != "" {
			if _, ok := seen[name]; !ok {
				myth = append(myth, name)
				seen[name] = struct{}{}
				fmt.Printf("Found myth element: %s\n", name)
			}
		}
	})

	log.Printf("Myth elements scraped (unique): %d\n", len(myth))
	return myth
}

func checkArray(elements []string, target string) bool {
	for _, el := range elements {
		if target == el {
			return true
		}
	}
	return false
}

func main() {
	mythsElements := getMythElements()
	fmt.Println(mythsElements)
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

	output := make(map[string]map[string][][]string)
	currentTierKey := ""
	totalElements := 0

	doc.Find("h2, h3, tr").Each(func(i int, s *goquery.Selection) {
		tag := goquery.NodeName(s)

		if tag == "h2" || tag == "h3" {
			heading := strings.TrimSpace(s.Text())
			if strings.HasPrefix(heading, "Tier") {
				parts := strings.Fields(heading)
				if len(parts) >= 2 {
					tierNum, err := strconv.Atoi(parts[1])
					if err == nil {
						currentTierKey = fmt.Sprintf("tier %d", tierNum)
						if _, ok := output[currentTierKey]; !ok {
							output[currentTierKey] = make(map[string][][]string)
						}
					}
				}
			}
			return
		}

		if tag != "tr" || currentTierKey == "" {
			return
		}

		left := s.Find("td").First()
		right := left.Next()

		element := strings.TrimSpace(left.Find("a").Text())
		if element == "" {
			return
		}

		if checkArray(mythsElements, element) {
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
				if checkArray(mythsElements, parts[0]) || checkArray(mythsElements, parts[1]) {
					return
				}
				recipes = append(recipes, parts)
			}
		})

		if len(recipes) > 0 {
			output[currentTierKey][element] = recipes
			totalElements++
		}
	})

	file, err := os.Create("data.json")
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

	fmt.Printf("Scraping completed. Data saved to data.json\n")
	fmt.Printf("Total elements with recipes: %d\n", totalElements)
}
