/* Kelompok   : Kelompok 21 - Deadliner Tobat                                    */
/* Nama - 1   : Muhammad Raihan Nazhim Oktana                                    */
/* NIM - 1    : K01 - 13523021 - Teknik Informatika (IF-Ganesha) ITB             */
/* Nama - 2   : Mayla Yaffa Ludmilla                                             */
/* NIM - 2    : K01 - 13523050 - Teknik Informatika (IF-Ganesha) ITB             */
/* Nama - 3   : Anella Utari Gunadi                                              */
/* NIM - 3    : K02 - 13523078 - Teknik Informatika (IF-Ganesha) ITB             */
/* Tanggal    : Senin, 12 Mei 2025                                               */
/* Tugas      : Tugas Besar 2 - Strategi Algoritma (IF2211-24)                   */
/* File Path  : Tubes2_DeadlinerTobat/src/backend/main_algorithm.go              */
/* Deskripsi  : F02 - Main Algorithm (BFS - DFS - BDR | Terminal CLI & Website)  */
/* PIC F02    : K01 - 13523021 - Muhammad Raihan Nazhim Oktana                   */

package backend;

import (
	"os";
	"fmt";
	"log";
	"time";
	"flag";
	"bufio";
	"unicode";
	"strconv";
	"strings";
	"sync/atomic";
)

var counter int64;

func touch() {
	atomic.AddInt64(&counter , 1);
}

func CustomizeCase(str string) string {
    if (str == "") {
        return str;
    } else {
		runes := []rune(str);
		runes[0] = unicode.ToUpper(runes[0]);
		for i := 1 ; i < len(runes) ; i++ {
			runes[i] = unicode.ToLower(runes[i]);
		}
		return string(runes);
	}
}

func DisplayTreeTerminal(node *RecipeNode, gallery *Gallery , depth int , max_tier int) {
    strh := strings.Repeat("-" , depth * 3) + ">";
    tier := 0;
    if element , check := gallery.GalleryName[node.Name] ; check {
        tier = element.Tier;
    }
    label := fmt.Sprintf("Tier %d" , tier);
	if (tier == 0 && tier == max_tier) {
		label = "Base Element & Target Element - Tier 0";
	} else if (tier == 0) {
		label = "Base Element - Tier 0";
	} else if (tier == max_tier) {
		label = fmt.Sprintf("Target Element - Tier %d" , max_tier);
	}
    fmt.Printf("%s \"%s\" (%s)\n" , strh , node.Name , label);
    for _ , child := range node.Parents {
        DisplayTreeTerminal(child , gallery , depth + 1 , max_tier);
    }
}

func DisplayResultTerminal(res AlgorithmResult , t time.Time , gallery *Gallery , target string , algorithm string , max_recipe *int) {
	ms := time.Since(t).Milliseconds();
	fmt.Println();
	fmt.Printf("Total Recipe : %d Recipe\n" , len(res.Trees));
	fmt.Printf("Visited Node : %d Node\n" , res.VisitedCount);
	fmt.Printf("Time Usage   : %d ms\n" , ms);
	fmt.Println();
	var answer string;
	mrp := len(res.Trees);
	for {
		if (*max_recipe != 0 && *max_recipe <= len(res.Trees)) {
			break;
		} else {
			fmt.Print("Apakah Anda yakin ingin melihat semua recipe? (Y/N) : ");
			scanner := bufio.NewReader(os.Stdin);
			input , _ := scanner.ReadString('\n');
			answer = strings.ToLower(strings.TrimSpace(input));
			if (answer != "y" && answer != "n") {
				fmt.Println("Input tidak valid, silakan coba lagi.");
			} else {
				break;
			}
		}
	}
	if (answer == "n") {
		*max_recipe = len(res.Trees);
		for {
			fmt.Printf("Output Recipe (0 - %d | Jika 0 Maka Tidak Ada Output) : " , *max_recipe);
			scanner := bufio.NewReader(os.Stdin);
			mrc , _ := scanner.ReadString('\n');
			if value , err := strconv.Atoi(strings.TrimSpace(mrc)) ; err == nil && 0 <= value && value <= *max_recipe {
				*max_recipe = value;
				break;
			}
			fmt.Printf("Input output recipe harus berupa bilangan bulat non-negatif dalam rentang (0 - %d).\n" , *max_recipe);
		}
		mrp = *max_recipe;
	} else {
		*max_recipe = len(res.Trees);
	}
	fmt.Println("==============================================================");
	fmt.Println();
	for i , node := range res.Trees {
		if (i >= mrp) {
			break;
		} else {
			fmt.Printf("=== Recipe #%d ===\n" , i + 1);
			DisplayTreeTerminal(node , gallery , 0 , GetTier(gallery , target));
			fmt.Println();
		}
    }
	fmt.Println("======= STATS & INPUT HISTORY =======");
	fmt.Printf("Target Element                 : %s\n" , target);
	fmt.Printf("Algorithm (BFS/DFS/BDR)        : %s\n" , algorithm);
	fmt.Printf("Display Recipe (0 = No Recipe) : %d\n" , *max_recipe);
	fmt.Println();
	fmt.Printf("Total Recipe : %d Recipe\n" , len(res.Trees));
	fmt.Printf("Visited Node : %d Node\n" , res.VisitedCount);
	fmt.Printf("Time Usage   : %d ms\n" , ms);
	fmt.Println();
}

func MainTerminal() {
	path := flag.String("data" , "data/data.json" , "JSON Location");
	flag.Parse();
	gallery , err := LoadRecipeGallery(*path);
	if (err != nil) {
		log.Fatal(err);
	} else {
		for {
			input := bufio.NewReader(os.Stdin);
			var target string;
			for {
				fmt.Print("Target Element              : ");
				tgt , _ := input.ReadString('\n');
				target = CustomizeCase(strings.TrimSpace(tgt));
				if _ , check := gallery.GalleryName[target] ; check || IsBase(target) {
					break;
				}
				fmt.Println("Input target elemen tidak valid.");
			}
			var algorithm string;
			for {
				fmt.Print("Algorithm (BFS/DFS/BDR)     : ");
				alg , _ := input.ReadString('\n');
				algorithm = strings.ToUpper(strings.TrimSpace(alg));
				if (algorithm == "BFS" || algorithm == "DFS" || algorithm == "BDR") {
					break;
				} else {
					fmt.Println("Input pilihan algoritma tidak valid.");
				}
			}
			var max_recipe int;
			for {
				fmt.Print("Max Recipe (0 = All Recipe) : ");
				mrc , _ := input.ReadString('\n');
				if value , err := strconv.Atoi(strings.TrimSpace(mrc)) ; err == nil && value >= 0 {
					max_recipe = value;
					break;
				}
				fmt.Println("Input max recipe harus berupa bilangan bulat non-negatif.");
			}
			atomic.StoreInt64(&counter , 0);
			start := time.Now();
			var res AlgorithmResult;
			EnableMultithreading();
			if (algorithm == "BFS") {
				res = BFS(gallery , target , max_recipe);
			} else if (algorithm == "DFS") {
				res = DFS(gallery , target , max_recipe);
			} else {
				res = BDR(gallery , target , max_recipe);
			}
			DisableMultithreading();
			DisplayResultTerminal(res , start , gallery , target , algorithm , &max_recipe);
			fmt.Println("==============================================================");
			var answer string;
			for {
				fmt.Print("Apakah ingin menjalankan program lagi? (Y/N) : ");
				scanner := bufio.NewReader(os.Stdin);
				input , _ := scanner.ReadString('\n');
				answer = strings.ToLower(strings.TrimSpace(input));
				if (answer != "y" && answer != "n") {
					fmt.Println("Input tidak valid, silakan coba lagi.");
				} else {
					fmt.Println("==============================================================");
					break;
				}
			}
			if (answer == "n") {
				break;
			}
		}
	}
}