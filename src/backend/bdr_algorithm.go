/* Kelompok   : Kelompok 21 - Deadliner Tobat                                    */
/* Nama - 1   : Muhammad Raihan Nazhim Oktana                                    */
/* NIM - 1    : K01 - 13523021 - Teknik Informatika (IF-Ganesha) ITB             */
/* Nama - 2   : Mayla Yaffa Ludmilla                                             */
/* NIM - 2    : K01 - 13523050 - Teknik Informatika (IF-Ganesha) ITB             */
/* Nama - 3   : Anella Utari Gunadi                                              */
/* NIM - 3    : K02 - 13523078 - Teknik Informatika (IF-Ganesha) ITB             */
/* Tanggal    : Senin, 12 Mei 2025                                               */
/* Tugas      : Tugas Besar 2 - Strategi Algoritma (IF2211-24)                   */
/* File Path  : src/backend/bdr_algorithm.go                                     */
/* Deskripsi  : B01 - Bidirectional Algorithm (Enumeration - Meet In The Middle) */
/* PIC B01    : K01 - 13523021 - Muhammad Raihan Nazhim Oktana                   */

package backend;

import (
    "sync";
    "sync/atomic";
)

func EnumerateTopBDR(gallery *Gallery , name string , mid_tier int , memory map[string][]*RecipeNode) []*RecipeNode {
    touch();
    tier := GetTier(gallery , name);
    if (tier <= mid_tier) {
        return []*RecipeNode{{Name : name}};
    } else if res , check := memory[name] ; check {
        return res;
    } else {
		var res []*RecipeNode;
		element := gallery.GalleryName[name];
		for _ , parent := range element.Parents {
			touch();
			l := parent[0];
			r := parent[1];
			if (GetTier(gallery , l) < tier && GetTier(gallery , r) < tier) {
				pl := EnumerateTopBDR(gallery , l , mid_tier , memory);
				pr := EnumerateTopBDR(gallery , r , mid_tier , memory);
				for _ , ll := range pl {
					for _ , rr := range pr {
						res = append(res , &RecipeNode{Name : name , Parents : []*RecipeNode{ll , rr}});
					}
				}
			}
		}
		memory[name] = res;
		return res;
	}
}

func EnumerateBotBDR(gallery *Gallery, name string, memory map[string][]*RecipeNode) []*RecipeNode {
    touch();
    if res , check := memory[name] ; check {
        return res;
    } else {
		element := gallery.GalleryName[name];
		if (element == nil || element.Tier == 0 || len(element.Parents) == 0) {
			res := []*RecipeNode{{Name : name}};
			memory[name] = res;
			return res;
		} else {
			var res []*RecipeNode
			for _ , parent := range element.Parents {
				touch();
				l := parent[0];
				r := parent[1];
				if (GetTier(gallery , l) < element.Tier && GetTier(gallery , r) < element.Tier) {
					pl := EnumerateBotBDR(gallery , l , memory);
					pr := EnumerateBotBDR(gallery , r , memory);
					for _ , l := range pl {
						for _ , r := range pr {
							res = append(res , &RecipeNode{Name : name , Parents : []*RecipeNode{l , r}});
						}
					}
				}
			}
			memory[name] = res;
			return res;
		}
	}
}

func BDR(gallery *Gallery, target string, max_recipe int) AlgorithmResult {
    if (GetTier(gallery , target) == 0) {
        touch();
        node := &RecipeNode{Name : target};
        return AlgorithmResult{Trees : []*RecipeNode{node} , VisitedCount : int(atomic.LoadInt64(&counter))};
    } else {
		if (max_recipe == 0) {
			max_recipe = int(^uint(0) >> 1);
		}
		mid_tier := GetMidTier(gallery , target);
		memory_top := make(map[string][]*RecipeNode);
		EnumerateTopBDR(gallery , target , mid_tier , memory_top);
		mnext := memory_top[target];
		group := make(map[string][]*RecipeNode);
		for _ , next := range mnext {
			var Collecting func(*RecipeNode);
			Collecting = func(node *RecipeNode) {
				if (len(node.Parents) == 0) {
					if (GetTier(gallery, node.Name) >= mid_tier) {
						group[node.Name] = append(group[node.Name] , next);
					}
				} else {
					for _ , parent := range node.Parents {
						Collecting(parent);
					}
				}
			}
			Collecting(next)
		}
		var (
			wg    sync.WaitGroup;
			res   []*RecipeNode;
			mutex sync.Mutex;
			signature_tree = make(map[string]struct{});
		)
		for name , parts := range group {
			wg.Add(1);
			go func() {
				defer wg.Done();
				memory_bot := make(map[string][]*RecipeNode);
				trees_bot := EnumerateBotBDR(gallery , name , memory_bot);
				var local_res []*RecipeNode;
				for _ , pt := range parts {
					for _ , bt := range trees_bot {
						touch();
						clone_top , clone_map := CloneTreeMap(pt);
						var leaf *RecipeNode;
						var FindLeaf func(*RecipeNode)
						FindLeaf = func(node *RecipeNode) {
							if (leaf != nil) {
								return;
							} else if (node.Name == name && len(node.Parents) == 0) {
								leaf = node;
							} else {
								for _ , parent := range node.Parents {
									FindLeaf(parent);
								}
							}
						}
						FindLeaf(pt);
						clone_bot , _ := CloneTreeMap(bt);
						clone_map[leaf].Parents = clone_bot.Parents;
						signature := SignatureTree(clone_top);
						mutex.Lock();
						if _ , check := signature_tree[signature] ; !check {
							signature_tree[signature] = struct{}{};
							local_res = append(local_res , clone_top);
							if (len(local_res) >= max_recipe) {
								mutex.Unlock();
								goto flush;
							}
						}
						mutex.Unlock();
					}
				}
				flush :
					mutex.Lock();
					res = append(res , local_res...);
					mutex.Unlock();
			}();
		}
		wg.Wait();
		if (len(res) > max_recipe) {
			res = res[:max_recipe];
		}
		return AlgorithmResult{Trees : res , VisitedCount : int(atomic.LoadInt64(&counter))};
	}
}