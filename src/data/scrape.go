/* Kelompok   : Kelompok 21 - Deadliner Tobat                          */
/* Nama - 1   : Muhammad Raihan Nazhim Oktana                          */
/* NIM - 1    : K01 - 13523021 - Teknik Informatika (IF-Ganesha) ITB   */
/* Nama - 2   : Mayla Yaffa Ludmilla                                   */
/* NIM - 2    : K01 - 13523050 - Teknik Informatika (IF-Ganesha) ITB   */
/* Nama - 3   : Anella Utari Gunadi                                    */
/* NIM - 3    : K02 - 13523078 - Teknik Informatika (IF-Ganesha) ITB   */
/* Tanggal    : Minggu, 11 Mei 2025                                    */
/* Tugas      : Tugas Besar 2 - Strategi Algoritma (IF2211-24)         */
/* File Path  : Tubes2_DeadlinerTobat/src/data/scrape.go               */
/* Deskripsi  : F06 - Scraping (Website Little Alchemy 2)              */
/* PIC F06    : K01 - 13523050 - Mayla Yaffa Ludmilla                  */

package main

import (
	"os"
	"fmt"
	"log"
	"strconv"
	"strings"
	"net/http"
	"encoding/json"
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

	// map[tier number]map[element][][]string
	output := make(map[int]map[string][][]string)
	currentTier := -1
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
						currentTier = tierNum
						if _, exists := output[currentTier]; !exists {
							output[currentTier] = make(map[string][][]string)
						}
					}
				}
			}
			return
		}

		if tag != "tr" || currentTier == -1 {
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
