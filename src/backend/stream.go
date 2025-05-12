package backend

func BFSStream(gallery *Gallery, target string, maxRecipe int, out chan<- *RecipeNode) {
    res := BFS(gallery, target, maxRecipe)

    for _, tree := range res.Trees {
        out <- tree // kirim 1 tree utuh per sekali kirim
    }
}