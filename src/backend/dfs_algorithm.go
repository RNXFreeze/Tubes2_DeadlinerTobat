/* Kelompok   : Kelompok 21 - Deadliner Tobat                          */
/* Nama - 1   : Muhammad Raihan Nazhim Oktana                          */
/* NIM - 1    : K01 - 13523021 - Teknik Informatika (IF-Ganesha) ITB   */
/* Nama - 2   : Mayla Yaffa Ludmilla                                   */
/* NIM - 2    : K01 - 13523050 - Teknik Informatika (IF-Ganesha) ITB   */
/* Nama - 3   : Anella Utari Gunadi                                    */
/* NIM - 3    : K02 - 13523078 - Teknik Informatika (IF-Ganesha) ITB   */
/* Tanggal    : Minggu, 11 Mei 2025                                    */
/* Tugas      : Tugas Besar 2 - Strategi Algoritma (IF2211-24)         */
/* File Path  : Tubes2_DeadlinerTobat/src/backend/dfs_algorithm.go     */
/* Deskripsi  : F04 - Depth First Search Algorithm (DFS)               */
/* PIC F04    : K01 - 13523021 - Muhammad Raihan Nazhim Oktana         */

package backend;

import (
	"sync/atomic"
)

type DFSOptions struct {
	MaxRecipes int;
	LiveChan   chan<- *RecipeNode;
}

type DFSResult struct {
	Trees        []*RecipeNode;
	VisitedCount int;
}

func CloneNode(node *RecipeNode) *RecipeNode {
	if (node == nil) {
		return nil;
	} else {
		parents := make([]*RecipeNode , len(node.Parents));
		for idx , cp := range node.Parents {
			parents[idx] = CloneNode(cp);
		}
		return &RecipeNode{Name : node.Name , Parents : parents};
	}
}

func CloneSlice(src []*RecipeNode) []*RecipeNode {
	cp := make([]*RecipeNode , len(src));
	for idx , node := range src {
		cp[idx] = CloneNode(node);
	}
	return cp;
}

func (res DFSResult) ToMultipleResult() BFSResult {
	return BFSResult(res);
}

func DFS(gallery *Gallery , target string , option DFSOptions) DFSResult {
	limit := option.MaxRecipes;
	if (limit == 0) {
		limit = int(^uint(0) >> 1);
	}
	var counter int64 = 0;
	stack := map[string]bool{};
	memory := map[string][]*RecipeNode{};
	visited := map[string]struct{}{};
	GetTier := func(n string) int {
		if element , check := gallery.GalleryName[n] ; check {
			return element.Tier;
		} else {
			return 0;
		}
	}
	var enumerate func(string) []*RecipeNode;
	enumerate = func(name string) []*RecipeNode {
		atomic.AddInt64(&counter , 1);
		if (stack[name]) {
			return []*RecipeNode{{Name : name}};
		} else if v , check := memory[name] ; check {
			return CloneSlice(v);
		}
		visited[name] = struct{}{};
		element := gallery.GalleryName[name];
		if (element == nil || len(element.Parents) == 0) {
			memory[name] = []*RecipeNode{{Name : name}};
			return memory[name];
		}
		stack[name] = true;
		var res []*RecipeNode;
		outer :
			for _ , r := range element.Parents {
				if (GetTier(r[0]) >= element.Tier || GetTier(r[1]) >= element.Tier) {
					continue;
				} else {
					left := enumerate(r[0]);
					right := enumerate(r[1]);
					for _ , ll := range left {
						for _ , rr := range right {
							res = append(res , &RecipeNode {
								Name    : name,
								Parents : []*RecipeNode{ll , rr},
							});
							if (len(res) >= limit) {
								break outer;
							}
						}
					}
				}
			}
			memory[name] = res;
			stack[name] = false;
			return CloneSlice(res);
	}
	tree := enumerate(target);
	if (len(tree) > limit) {
		tree = tree[:limit];
	}
	if (option.LiveChan != nil) {
		go func() {
			for _ , t := range tree {
				option.LiveChan <- t;
			}
		}();
	}
	return DFSResult{Trees : tree , VisitedCount : int(counter)};
}