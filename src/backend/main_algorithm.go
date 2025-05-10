/* Kelompok   : Kelompok 21 - Deadliner Tobat                          */
/* Nama - 1   : Muhammad Raihan Nazhim Oktana                          */
/* NIM - 1    : K01 - 13523021 - Teknik Informatika (IF-Ganesha) ITB   */
/* Nama - 2   : Mayla Yaffa Ludmilla                                   */
/* NIM - 2    : K01 - 13523050 - Teknik Informatika (IF-Ganesha) ITB   */
/* Nama - 3   : Anella Utari Gunadi                                    */
/* NIM - 3    : K02 - 13523078 - Teknik Informatika (IF-Ganesha) ITB   */
/* Tanggal    : Minggu, 11 Mei 2025                                    */
/* Tugas      : Tugas Besar 2 - Strategi Algoritma (IF2211-24)         */
/* File Path  : Tubes2_DeadlinerTobat/src/backend/main_algorithm.go    */
/* Deskripsi  : F02 - Main Algorithm (BFS - DFS - BIDERECTIONAL)       */
/* PIC F02    : K01 - 13523021 - Muhammad Raihan Nazhim Oktana         */

package backend

import (
	"os";
	"fmt";
	"log";
	"time";
	"flag";
	"bufio";
	"strconv";
	"strings";
	"encoding/json";
)

func DisplayResultTerminal(res BFSMultiResult , t time.Time) {
	ms := time.Since(t).Milliseconds();
	fmt.Println();
	fmt.Printf("Total Recipe : %d Recipe\n" , len(res.Trees));
	fmt.Printf("Visited Node : %d Node\n" , res.VisitedCount);
	fmt.Printf("Time Usage   : %d ms\n" , ms);
	fmt.Println();
	for i , t := range res.Trees {
		b , _ := json.MarshalIndent(t , "" , "  ");
		fmt.Printf("--- Recipe #%d ---\n%s\n\n" , i + 1 , string(b));
	}
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
			fmt.Print("Target Element              : ");
			target , _ := input.ReadString('\n');
			target = strings.TrimSpace(target)
			fmt.Print("Algorithm (BFS/DFS/BDR)     : ");
			algorithm , _ := input.ReadString('\n');
			algorithm = strings.ToUpper(strings.TrimSpace(algorithm));
			fmt.Print("Max Recipe (0 = All Recipe) : ");
			max_input , _ := input.ReadString('\n');
			max_recipe , _ := strconv.Atoi(strings.TrimSpace(max_input));
			start := time.Now();
			if (algorithm == "BFS") {
				res := BFS(gallery , target , max_recipe);
				DisplayResultTerminal(res , start);
			} else if (algorithm == "DFS") {
				res := DFS(gallery , target , DFSOptions{MaxRecipes : max_recipe});
				DisplayResultTerminal(res.toMulti() , start);
			} else if (algorithm == "BDR") {
				res := BDR(gallery , target , max_recipe);
				DisplayResultTerminal(res , start);
			} else {
				fmt.Println("Input pilihan algoritma tidak valid.");
			}
			var answer string;
			fmt.Println("=====================================================");
			for {
				fmt.Print("Apakah ingin menjalankan program lagi? (Y/N) : ");
				scanner := bufio.NewReader(os.Stdin);
				input , _ := scanner.ReadString('\n');
				answer = strings.ToLower(strings.TrimSpace(input));
				if (answer != "y" && answer != "n") {
					fmt.Println("Input tidak valid, silakan coba lagi.");
				}
				if (answer == "y") {
					fmt.Println("=====================================================");
					break;
				}
			}
			if (answer == "n") {
				break;
			}
		}
	}
}

func MainWebsite() {

}
