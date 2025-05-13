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
    "sync";
    "sync/atomic";
)

func DFS(gallery *Gallery , target string , max_recipe int) AlgorithmResult {
    if (max_recipe == 0) {
        max_recipe = int(^uint(0) >> 1);
    }
    element := gallery.GalleryName[target];
    if (element == nil || element.Tier == 0 || len(element.Parents) == 0) {
        touch();
        node := &RecipeNode{Name : target};
        return AlgorithmResult{Trees : []*RecipeNode{node} , VisitedCount : int(atomic.LoadInt64(&counter))};
    } else {
		var (
			res    []*RecipeNode;
			mutex  sync.Mutex;
			memory sync.Map;
		)
		var EnumerateDFS func(string , map[string]bool) []*RecipeNode;
		EnumerateDFS = func(name string , stack map[string]bool) []*RecipeNode {
			touch();
			if (stack[name]) {
				return []*RecipeNode{{Name : name}};
			} else if value , check := memory.Load(name) ; check {
				return CloneSlice(value.([]*RecipeNode));
			} else {
				element := gallery.GalleryName[name];
				if (element == nil || element.Tier == 0 || len(element.Parents) == 0) {
					base := []*RecipeNode{{Name : name}};
					memory.Store(name , base);
					return base;
				} else {
					stack[name] = true;
					var local_res []*RecipeNode;
					for _ , parent := range element.Parents {
						l := parent[0];
						r := parent[1];
						if (GetTier(gallery , l) < element.Tier && GetTier(gallery , r) < element.Tier) {
							pl := EnumerateDFS(l , stack);
							pr := EnumerateDFS(r , stack);
							for _ , ll := range pl {
								for _ , rr := range pr {
									touch();
									local_res = append(local_res , &RecipeNode{Name: name , Parents: []*RecipeNode{ll , rr}});
									if (len(local_res) >= max_recipe) {
										memory.Store(name , local_res);
										stack[name] = false;
										return CloneSlice(local_res);
									}
								}
							}
						}
					}
					memory.Store(name , local_res);
					stack[name] = false;
					return CloneSlice(local_res);
				}
			}
		}
		var wg sync.WaitGroup;
		for _ , parent := range element.Parents {
			l := parent[0];
			r := parent[1];
			if (GetTier(gallery , l) < element.Tier && GetTier(gallery , r) < element.Tier) {
				wg.Add(1);
				go func(l string , r string) {
					defer wg.Done();
					local_stack := map[string]bool{};
					pl := EnumerateDFS(l , local_stack);
					pr := EnumerateDFS(r , local_stack);
					var local_res []*RecipeNode;
					for _ , ll := range pl {
						for _ , rr := range pr {
							touch();
							local_res = append(local_res , &RecipeNode{Name : target , Parents : []*RecipeNode{ll , rr}});
							if (len(local_res) >= max_recipe) {
								break;
							}
						}
						if (len(local_res) >= max_recipe) {
							break;
						}
					}
					mutex.Lock();
					res = append(res , local_res...);
					mutex.Unlock();
				}(l , r);
			}
		}
		wg.Wait();
		if (len(res) > max_recipe) {
			res = res[:max_recipe];
		}
		return AlgorithmResult{Trees : res , VisitedCount : int(atomic.LoadInt64(&counter))};
	}
}