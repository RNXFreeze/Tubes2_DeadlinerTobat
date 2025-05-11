/* Kelompok   : Kelompok 21 - Deadliner Tobat                          */
/* Nama - 1   : Muhammad Raihan Nazhim Oktana                          */
/* NIM - 1    : K01 - 13523021 - Teknik Informatika (IF-Ganesha) ITB   */
/* Nama - 2   : Mayla Yaffa Ludmilla                                   */
/* NIM - 2    : K01 - 13523050 - Teknik Informatika (IF-Ganesha) ITB   */
/* Nama - 3   : Anella Utari Gunadi                                    */
/* NIM - 3    : K02 - 13523078 - Teknik Informatika (IF-Ganesha) ITB   */
/* Tanggal    : Minggu, 11 Mei 2025                                    */
/* Tugas      : Tugas Besar 2 - Strategi Algoritma (IF2211-24)         */
/* File Path  : Tubes2_DeadlinerTobat/src/backend/bfs_algorithm.go     */
/* Deskripsi  : F03 - Breadth First Search Algorithm (BFS)             */
/* PIC F03    : K01 - 13523021 - Muhammad Raihan Nazhim Oktana         */

package backend;

import (
	"sync/atomic"
)

type BFSResult struct {
	Trees        []*RecipeNode;
	VisitedCount int;
}

func BFS(gallery *Gallery , target string , max_recipe int) BFSResult {
	if (max_recipe == 0) {
		max_recipe = int(^uint(0) >> 1);
	} else if (max_recipe < 0) {
		max_recipe = 1;
	}
	var counter int64 = 0;
	visited := map[string]struct{}{};
	GetTier := func(n string) int {
		if element , check := gallery.GalleryName[n] ; check {
			return element.Tier;
		}
		return 0;
	}
	element := gallery.GalleryName[target];
	if (element == nil || len(element.Parents) == 0) {
		atomic.AddInt64(&counter , 1);
		return BFSResult{Trees : nil , VisitedCount : int(counter)};
	} else {
		var result []*RecipeNode;
		for _ , rec := range element.Parents {
			if (GetTier(rec[0]) >= element.Tier || GetTier(rec[1]) >= element.Tier) {
				continue;
			} else {
				atomic.AddInt64(&counter , 1);
				root := &RecipeNode{Name : target}
				root.Parents = []*RecipeNode {
					BuildSubBFS(gallery , rec[0] , visited , GetTier , &counter),
					BuildSubBFS(gallery , rec[1] , visited , GetTier , &counter),
				}
				result = append(result , root);
				if (len(result) >= max_recipe) {
					break;
				}
			}
		}
		return BFSResult{Trees : result , VisitedCount : int(counter)};
	}
}

func BuildSubBFS(gallery *Gallery , name string , visited map[string]struct{} , GetTier func(string) int , counter *int64) *RecipeNode {
	root := &RecipeNode{Name : name};
	queue := []*RecipeNode{root};
	for (len(queue) > 0) {
		atomic.AddInt64(counter , 1);
		cur := queue[0];
		queue = queue[1:];
		visited[cur.Name] = struct{}{};
		element := gallery.GalleryName[cur.Name];
		if (element == nil || len(element.Parents) == 0) {
			continue;
		} else {
			for _ , r := range element.Parents {
				if (GetTier(r[0]) >= element.Tier || GetTier(r[1]) >= element.Tier) {
					continue;
				} else {
					ll := &RecipeNode{Name : r[0]};
					rr := &RecipeNode{Name : r[1]};
					cur.Parents = []*RecipeNode{ll ,  rr};
					queue = append(queue , ll , rr);
					break;
				}
			}
		}
	}
	return root;
}