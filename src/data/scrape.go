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
	"os";
	"fmt";
	"log";
	"sort";
	"strconv";
	"strings";
	"net/http";
	"encoding/json";
	"github.com/PuerkitoBio/goquery";
)

func getMythElements() []string {
	var myth []string
	seen := make(map[string]struct{})

	url := "https://little-alchemy.fandom.com/wiki/Elements_(Myths_and_Monsters)"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
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
		name := strings.TrimSpace(row.Find("td").Eq(0).Text())
		if name != "" {
			if _, ok := seen[name]; !ok {
				myth = append(myth, name)
				seen[name] = struct{}{}
			}
		}
	})

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
	// 1) Dapatkan elemen Myth/Monster untuk di-skip
	mythElements := getMythElements()

	// 2) Fetch halaman utama Little Alchemy 2
	client := &http.Client{}
	req, err := http.NewRequest("GET",
		"https://little-alchemy.fandom.com/wiki/Elements_(Little_Alchemy_2)", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 3) Siapkan struktur output dan inisialisasi tier 0 (base elements)
	output := make(map[string]map[string][][]string)
	output["tier 0"] = make(map[string][][]string)
	currentTier := "tier 0"
	totalElements := 0

	doc.Find("h2, h3, tr").Each(func(_ int, s *goquery.Selection) {
		tag := goquery.NodeName(s)

		// Deteksi heading Tier N → override currentTier
		if (tag == "h2" || tag == "h3") && strings.HasPrefix(strings.TrimSpace(s.Text()), "Tier") {
			parts := strings.Fields(s.Text())
			if len(parts) >= 2 {
				if tierNum, err := strconv.Atoi(parts[1]); err == nil {
					currentTier = fmt.Sprintf("tier %d", tierNum)
					if _, exists := output[currentTier]; !exists {
						output[currentTier] = make(map[string][][]string)
					}
				}
			}
			return
		}

		// Proses hanya baris <tr> jika currentTier sudah ada
		if tag != "tr" || currentTier == "" {
			return
		}

		// Nama elemen
		element := strings.TrimSpace(s.Find("td").First().Find("a").Text())
		if element == "" || checkArray(mythElements, element) {
			return
		}

		// Tambahkan elemen ke output (jaga unique)
		if _, exists := output[currentTier][element]; !exists {
			output[currentTier][element] = nil
			totalElements++
		}

		// Kumpulkan resep jika ada
		var recipes [][]string
		s.Find("li").Each(func(_ int, li *goquery.Selection) {
			var pair []string
			li.Find("a").Each(func(_ int, a *goquery.Selection) {
				if txt := strings.TrimSpace(a.Text()); txt != "" {
					pair = append(pair, txt)
				}
			})
			if len(pair) == 2 &&
				!checkArray(mythElements, pair[0]) &&
				!checkArray(mythElements, pair[1]) {
				recipes = append(recipes, []string{pair[0], pair[1]})
			}
		})
		if len(recipes) > 0 {
			output[currentTier][element] = recipes
		}
	})

	// 4) Tuliskan JSON dengan urutan tier 0..15
	file, err := os.Create("data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	tiers := make([]string, 0, len(output))
	for t := range output {
		tiers = append(tiers, t)
	}
	sort.Slice(tiers, func(i, j int) bool {
		ni, _ := strconv.Atoi(strings.TrimPrefix(tiers[i], "tier "))
		nj, _ := strconv.Atoi(strings.TrimPrefix(tiers[j], "tier "))
		return ni < nj
	})

	file.WriteString("{\n")
	for idx, tier := range tiers {
		b, err := json.MarshalIndent(output[tier], "  ", "  ")
		if err != nil {
			log.Fatalf("marshal tier %s: %v", tier, err)
		}
		file.WriteString(fmt.Sprintf("  %q: %s", tier, b))
		if idx < len(tiers)-1 {
			file.WriteString(",\n")
		} else {
			file.WriteString("\n")
		}
	}
	file.WriteString("}\n")

	fmt.Printf("Scraping completed. Total elements: %d\n", totalElements)
}