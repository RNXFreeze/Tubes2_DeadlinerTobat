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

package backend

import "sync/atomic"

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

func BDR(gallery *Gallery , target string , option AlgorithmOption) AlgorithmResult {
    if (GetTier(gallery , target) == 0) {
        node := &RecipeNode{Name : target}
        return AlgorithmResult{Trees : []*RecipeNode{node} , VisitedCount : 1};
    } else {
		max_recipe := option.MaxRecipes;
		if (max_recipe == 0) {
			max_recipe = int(^uint(0) >> 1);
		}
		mid_tier := GetMidTier(gallery , target);
		memory_top := make(map[string][]*RecipeNode);
		EnumerateTopBDR(gallery , target , mid_tier , memory_top);
		nxt := memory_top[target];
		grp := make(map[string][]*RecipeNode);
		for _ , nt := range nxt {
			var Collecting func(*RecipeNode);
			Collecting = func(node *RecipeNode) {
				if (len(node.Parents) == 0) {
					if (GetTier(gallery , node.Name) >= mid_tier) {
						grp[node.Name] = append(grp[node.Name] , nt)
					}
				} else {
					for _ , nodes := range node.Parents {
						Collecting(nodes);
					}
				}
			}
			Collecting(nt);
		}
		var res []*RecipeNode;
		signature_tree := make(map[string]struct{});
		memory_bot := make(map[string][]*RecipeNode);
		for name , part := range grp {
			trees_bot := EnumerateBotBDR(gallery , name , memory_bot);
			for _ , pt := range part {
				for _ , bt := range trees_bot {
					touch();
					clone_top , clone_map := CloneTreeMap(pt)
					var leaf *RecipeNode;
					var Finding func(*RecipeNode);
					Finding = func(node *RecipeNode) {
						if (leaf == nil) {
							if (node.Name == name && len(node.Parents) == 0) {
								leaf = node;
							} else {
								for _ , nodes := range node.Parents {
									Finding(nodes);
								}
							}
						}
					}
					Finding(pt);
					clone_bot , _ := CloneTreeMap(bt);
					clone_map[leaf].Parents = clone_bot.Parents;
					signature := SignatureTree(clone_top);
					if _ , check := signature_tree[signature] ; !check {
						signature_tree[signature] = struct{}{};
						res = append(res , clone_top);
						if (len(res) >= max_recipe) {
							return AlgorithmResult{Trees : res , VisitedCount : int(atomic.LoadInt64(&counter))};
						}
					}
				}
			}
		}
		if (len(res) > max_recipe) {
			res = res[:max_recipe];
		}
		if (option.LiveChan != nil) {
			go func() {
				for _ , t := range res {
					option.LiveChan <- t;
				}
			}();
		}
		return AlgorithmResult{Trees : res , VisitedCount : int(atomic.LoadInt64(&counter))};
	}
}