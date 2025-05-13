/* Kelompok   : Kelompok 21 - Deadliner Tobat                                    */
/* Nama - 1   : Muhammad Raihan Nazhim Oktana                                    */
/* NIM - 1    : K01 - 13523021 - Teknik Informatika (IF-Ganesha) ITB             */
/* Nama - 2   : Mayla Yaffa Ludmilla                                             */
/* NIM - 2    : K01 - 13523050 - Teknik Informatika (IF-Ganesha) ITB             */
/* Nama - 3   : Anella Utari Gunadi                                              */
/* NIM - 3    : K02 - 13523078 - Teknik Informatika (IF-Ganesha) ITB             */
/* Tanggal    : Senin, 12 Mei 2025                                               */
/* Tugas      : Tugas Besar 2 - Strategi Algoritma (IF2211-24)                   */
/* File Path  : Tubes2_DeadlinerTobat/src/backend/stream.go                      */
/* Deskripsi  : F00B - Stream API Helper (Live Status Update Tree)               */
/* PIC F00B   : K01 - 13523050 - Mayla Yaffa Ludmilla                            */

package backend

func BFSStream(gallery *Gallery, target string, maxRecipe int, out chan<- *RecipeNode) {
    res := BFS(gallery, target, AlgorithmOption{MaxRecipes: maxRecipe, LiveChan: out})

    for _, tree := range res.Trees {
        out <- tree // kirim 1 tree utuh per sekali kirim
    }
}

func BFSStream1(gallery *Gallery , target string , option AlgorithmOption) {
	visited := make(map[string]bool);
	queue := []*RecipeNode{};
	res := BFS(gallery , target , option);
	queue = append(queue , res.Trees...);
	out := option.LiveChan;
	for (len(queue) > 0) {
		cur := queue[0];
		queue = queue[1:];
		if (!visited[cur.Name]) {
			visited[cur.Name] = true;
			out <- cur;
			queue = append(queue , cur.Parents...);
		}
	}
}

func DFSStream(gallery *Gallery , target string , option AlgorithmOption) {
	visited := make(map[string]bool);
	queue := []*RecipeNode{};
	res := DFS(gallery , target , option);
	queue = append(queue , res.Trees...);
	out := option.LiveChan;
	for (len(queue) > 0) {
		cur := queue[0];
		queue = queue[1:];
		if (!visited[cur.Name]) {
			visited[cur.Name] = true;
			out <- cur;
			queue = append(queue , cur.Parents...);
		}
	}
}

func BDRStream(gallery *Gallery , target string , option AlgorithmOption) {
	visited := make(map[string]bool);
	queue := []*RecipeNode{};
	res := BDR(gallery , target , option);
	queue = append(queue , res.Trees...);
	out := option.LiveChan;
	for (len(queue) > 0) {
		cur := queue[0];
		queue = queue[1:];
		if (!visited[cur.Name]) {
			visited[cur.Name] = true;
			out <- cur;
			queue = append(queue , cur.Parents...);
		}
	}
}