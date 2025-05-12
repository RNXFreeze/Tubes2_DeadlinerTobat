package backend

func BFSStream(gallery *Gallery, target string, maxRecipe int, out chan<- *RecipeNode) {
	visited := make(map[string]bool)
	type queueItem struct {
		Node *RecipeNode
	}

	queue := []*RecipeNode{}
	res := BFS(gallery, target, maxRecipe)
	queue = append(queue, res.Trees...)

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if visited[cur.Name] {
			continue
		}
		visited[cur.Name] = true
		out <- cur

		for _, parent := range cur.Parents {
			queue = append(queue, parent)
		}
	}
}
