/* Kelompok   : Kelompok 21 - Deadliner Tobat                                    */
/* Nama - 1   : Muhammad Raihan Nazhim Oktana                                    */
/* NIM - 1    : K01 - 13523021 - Teknik Informatika (IF-Ganesha) ITB             */
/* Nama - 2   : Mayla Yaffa Ludmilla                                             */
/* NIM - 2    : K01 - 13523050 - Teknik Informatika (IF-Ganesha) ITB             */
/* Nama - 3   : Anella Utari Gunadi                                              */
/* NIM - 3    : K02 - 13523078 - Teknik Informatika (IF-Ganesha) ITB             */
/* Tanggal    : Senin, 12 Mei 2025                                               */
/* Tugas      : Tugas Besar 2 - Strategi Algoritma (IF2211-24)                   */
/* File Path  : Tubes2_DeadlinerTobat/src/backend/dfs_algorithm.go               */
/* Deskripsi  : F04 - Depth First Search Algorithm (DFS - Stack & Memoization)   */
/* PIC F04    : K01 - 13523021 - Muhammad Raihan Nazhim Oktana                   */

package backend;

import (
	"time";
	"sync/atomic";
)

func DFS(gallery *Gallery , target string , option AlgorithmOption) AlgorithmResult {
	max_recipe := option.MaxRecipes;
	if (max_recipe == 0) {
		max_recipe = int(^uint(0) >> 1);
	}
	stack := map[string]bool{};
	memory := map[string][]*RecipeNode{};
	var EnumerateDFS func(string) []*RecipeNode;
	EnumerateDFS = func(name string) []*RecipeNode {
		touch();
		if (stack[name]) {
			return []*RecipeNode{{Name : name}};
		} else if value , check := memory[name] ; check {
			return CloneSlice(value);
		} else {
			element := gallery.GalleryName[name];
			if (element == nil || element.Tier == 0 || len(element.Parents) == 0) {
				memory[name] = []*RecipeNode{{Name : name}};
				return memory[name];
			} else {
				stack[name] = true;
				var res []*RecipeNode;
				for _ , parent := range element.Parents {
					if (GetTier(gallery , parent[0]) < element.Tier && GetTier(gallery , parent[1]) < element.Tier) {
						pl := EnumerateDFS(parent[0]);
						pr := EnumerateDFS(parent[1]);
						for _ , ll := range pl {
							for _ , rr := range pr { 
								res = append(res , &RecipeNode {Name : name , Parents : []*RecipeNode{ll , rr}});
								if (len(res) >= max_recipe) {
									memory[name] = res;
									stack[name] = false;
									return CloneSlice(res);
								}
							}
						}
					}
				}
				memory[name] = res;
				stack[name] = false;
				return CloneSlice(res);
			}
		}
	}
	res := EnumerateDFS(target);
	if (len(res) > max_recipe) {
		res = res[:max_recipe];
	}
	if (option.LiveChan != nil) {
		go func() {
			for _ , t := range res {
				option.LiveChan <- t;
				time.Sleep(1500 * time.Millisecond);
			}
		}();
	}
	return AlgorithmResult{Trees : res , VisitedCount : int(atomic.LoadInt64(&counter))};
}