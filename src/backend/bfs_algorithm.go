/* Kelompok   : Kelompok 21 - Deadliner Tobat                                    */
/* Nama - 1   : Muhammad Raihan Nazhim Oktana                                    */
/* NIM - 1    : K01 - 13523021 - Teknik Informatika (IF-Ganesha) ITB             */
/* Nama - 2   : Mayla Yaffa Ludmilla                                             */
/* NIM - 2    : K01 - 13523050 - Teknik Informatika (IF-Ganesha) ITB             */
/* Nama - 3   : Anella Utari Gunadi                                              */
/* NIM - 3    : K02 - 13523078 - Teknik Informatika (IF-Ganesha) ITB             */
/* Tanggal    : Senin, 12 Mei 2025                                               */
/* Tugas      : Tugas Besar 2 - Strategi Algoritma (IF2211-24)                   */
/* File Path  : Tubes2_DeadlinerTobat/src/backend/bfs_algorithm.go               */
/* Deskripsi  : F03 - Breadth First Search Algorithm (BFS - Queue & Signature)   */
/* PIC F03    : K01 - 13523021 - Muhammad Raihan Nazhim Oktana                   */

package backend;

import (
    "sync";
    "sync/atomic";
)

func BFS(gallery *Gallery , target string , max_recipe int) AlgorithmResult {
    if max_recipe == 0 {
        max_recipe = int(^uint(0) >> 1)
    }
    touch();
    if element , check := gallery.GalleryName[target] ; (check && element.Tier == 0) {
        node := &RecipeNode{Name : target};
        return AlgorithmResult{Trees : []*RecipeNode{node} , VisitedCount : int(atomic.LoadInt64(&counter))};
    } else {
		parent_map := make(map[string][][2]string);
		GetParentPairs := func(name string) [][2]string {
			if parent_pairs , check := parent_map[name] ; check {
				return parent_pairs;
			} else {
				var all_parent_pairs [][2]string;
				if element , check := gallery.GalleryName[name] ; check {
					for _ , parent := range element.Parents {
						all_parent_pairs = append(all_parent_pairs , [2]string{parent[0] , parent[1]});
					}
				}
				parent_map[name] = all_parent_pairs;
				return all_parent_pairs;
			}
		}
		var (
			res   []*RecipeNode;
			next  []PartialTree;
			queue []PartialTree;
			mutex_res  sync.Mutex;
			mutex_next sync.Mutex;
			signature_tree = make(map[string]struct{});
		)
		queue = append(queue , PartialTree{tree : &RecipeNode{Name : target} , leaf : []*RecipeNode{{Name : target}}});
		for (len(queue) > 0 && len(res) < max_recipe) {
			var wg sync.WaitGroup;
			next = next[:0];
			for _ , cur := range queue {
				exp := cur.leaf[0];
				element := gallery.GalleryName[exp.Name];
				for _ , parent := range GetParentPairs(exp.Name) {
					touch();
					l := parent[0];
					r := parent[1];
					if (GetTier(gallery , l) < element.Tier && GetTier(gallery , r) < element.Tier) {
						wg.Add(1);
						go func(l string , r string , cur PartialTree) {
							defer wg.Done()
							new_root , _ := CloneTreeMap(cur.tree);
							ptr := FindPointer(new_root, cur.leaf[0].Name)
							if (ptr == nil) {
								return;
							} else {
								pl := &RecipeNode{Name : l};
								pr := &RecipeNode{Name : r};
								ptr.Parents = []*RecipeNode{pl , pr};
								var new_leaf []*RecipeNode;
								for _ , leaf := range cur.leaf[1:] {
									if np := FindPointer(new_root , leaf.Name) ; np != nil {
										new_leaf = append(new_leaf , np);
									}
								}
								if (IsExpandable(gallery.GalleryName[l])) {
									new_leaf = append(new_leaf , pl);
								}
								if (IsExpandable(gallery.GalleryName[r])) {
									new_leaf = append(new_leaf , pr);
								}
								if (len(new_leaf) == 0) {
									signature := SignatureTree(new_root);
									mutex_res.Lock();
									if _ , check := signature_tree[signature]; (!check && len(res) < max_recipe) {
										signature_tree[signature] = struct{}{};
										res = append(res , new_root);
									}
									mutex_res.Unlock();
								} else {
									mutex_next.Lock();
									next = append(next , PartialTree{tree : new_root , leaf : new_leaf});
									mutex_next.Unlock();
								}
							}
						}(l , r , cur);
					}
				}
			}
			wg.Wait();
			queue , next = next , queue;
		}
		if (len(res) > max_recipe) {
			res = res[:max_recipe];
		}
		return AlgorithmResult {Trees : res , VisitedCount : int(atomic.LoadInt64(&counter))};
	}
}