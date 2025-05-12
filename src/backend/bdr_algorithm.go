/* Kelompok   : Kelompok 21 - Deadliner Tobat                          */
/* Nama - 1   : Muhammad Raihan Nazhim Oktana                          */
/* NIM - 1    : K01 - 13523021 - Teknik Informatika (IF-Ganesha) ITB   */
/* Nama - 2   : Mayla Yaffa Ludmilla                                   */
/* NIM - 2    : K01 - 13523050 - Teknik Informatika (IF-Ganesha) ITB   */
/* Nama - 3   : Anella Utari Gunadi                                    */
/* NIM - 3    : K02 - 13523078 - Teknik Informatika (IF-Ganesha) ITB   */
/* Tanggal    : Minggu, 11 Mei 2025                                    */
/* Tugas      : Tugas Besar 2 - Strategi Algoritma (IF2211-24)         */
/* File Path  : Tubes2_DeadlinerTobat/src/backend/bdr_algorithm.go     */
/* Deskripsi  : B01 - Bidirectional Algorithm (BDR)                    */
/* PIC B01    : K01 - 13523021 - Muhammad Raihan Nazhim Oktana         */

package backend

import (
	"sync";
	"sync/atomic";
)

func GetElementTier(gallery *Gallery , name string) int {
	if element , check := gallery.GalleryName[name] ; check {
		return element.Tier;
	} else {
		return 0;
	}
}

func EnumerateBDR(gallery *Gallery , name string , limit int , memory *sync.Map , visited *sync.Map , stack map[string]bool , counter *int64) []*RecipeNode {
	atomic.AddInt64(counter , 1);
	if (limit == 0) {
		return nil;
	} else if value , check := memory.Load(name) ; check {
		return CloneSlice(value.([]*RecipeNode));
	} else {
		visited.LoadOrStore(name , struct{}{});
		element := gallery.GalleryName[name];
		if (element == nil || element.Tier == 0 || len(element.Parents) == 0 || stack[name]) {
			res := []*RecipeNode{{Name : name}};
			memory.Store(name , res);
			return CloneSlice(res);
		} else {
			stack[name] = true;
			var out []*RecipeNode;
			outer :
				for _ , rec := range element.Parents {
					a , b := rec[0] , rec[1];
					if (GetElementTier(gallery , a) >= element.Tier || GetElementTier(gallery , b) >= element.Tier) {
						continue;
					} else {
						el := EnumerateBDR(gallery , a , limit , memory , visited , stack , counter);
						er := EnumerateBDR(gallery , b , limit , memory , visited , stack , counter);
						for _ , l := range el {
							for _ , r := range er {
								out = append(out , &RecipeNode {Name : name , Parents : []*RecipeNode{l , r}});
								if (len(out) >= limit) {
									break outer;
								}
							}
						}
					}
				}
			stack[name] = false;
			memory.Store(name , out);
			return CloneSlice(out);
		}
	}
}

func BDR(gallery *Gallery , target string , max_recipe int) BFSResult {
	if (max_recipe == 0) {
		max_recipe = int(^uint(0) >> 1);
	} else if (max_recipe < 0) {
		max_recipe = 1;
	}
	root := gallery.GalleryName[target];
	if (root == nil || len(root.Parents) == 0) {
		leaf := &RecipeNode{Name : target};
		return BFSResult{Trees : []*RecipeNode{leaf} , VisitedCount : 1};
	} else {
		var (
			counter int64;
			memory  sync.Map;
			visited sync.Map;
			sync_wg sync.WaitGroup;
			channel = make(chan *RecipeNode , 32);
		)
		for _ , rec := range root.Parents {
			a , b := rec[0] , rec[1];
			if (GetElementTier(gallery , a) >= root.Tier || GetElementTier(gallery, b) >= root.Tier) {
				continue;
			} else {
				sync_wg.Add(1);
				go func(x , y string) {
					defer sync_wg.Done();
					el := EnumerateBDR(gallery , x , max_recipe , &memory , &visited , map[string]bool{} , &counter);
					er := EnumerateBDR(gallery , y , max_recipe , &memory , &visited , map[string]bool{} , &counter);
					for _ , l := range el {
						for _ , r := range er {
							channel <- &RecipeNode {Name : target , Parents : []*RecipeNode{l , r}};
						}
					}
				} (a , b)
			}
		}
		go func() {
			sync_wg.Wait();
			close(channel);
		}();
		var trees []*RecipeNode;
		collect :
			for t := range channel {
				trees = append(trees , t);
				if (len(trees) >= max_recipe) {
					go func() {
						for range channel{};
					}();
					break collect;
				}
			}
		visited.Range(func(_ any , _ any) bool {
			return true;
		});
		if (len(trees) == 0) {
			return BFS(gallery , target , max_recipe);
		} else {
			return BFSResult{Trees : trees , VisitedCount : int(counter)}
		}
	}
}