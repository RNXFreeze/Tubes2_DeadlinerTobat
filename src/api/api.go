package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"Tubes2_DeadlinerTobat/src/backend"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type bfsResponse struct {
	Target       string                `json:"target"`
	VisitedCount int                   `json:"visited_count"`
	Trees        []*backend.RecipeNode `json:"trees"`
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gallery, err := backend.LoadRecipeGallery("data/data.json")
	if err != nil {
		log.Fatalf("load gallery: %v", err)
	}

	r.GET("/api/bfs", func(c *gin.Context) {
		target := c.Query("target")
		if target == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "query param 'target' is required"})
			return
		}

		maxRecipe, _ := strconv.Atoi(c.DefaultQuery("max_recipe", "0"))
		res := backend.BFS(gallery, target, maxRecipe)

		c.JSON(http.StatusOK, bfsResponse{
			Target:       target,
			VisitedCount: res.VisitedCount,
			Trees:        res.Trees, // sudah JSON-marshal-able karena ada tag di RecipeNode
		})
	})
	r.GET("/api/dfs", func(c *gin.Context) {
		target := c.Query("target")
		if target == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "query param 'target' is required"})
			return
		}

		maxRecipe, _ := strconv.Atoi(c.DefaultQuery("max_recipe", "0"))

		res := backend.DFS(gallery, target, backend.DFSOptions{
			MaxRecipes: maxRecipe,
		})

		c.JSON(http.StatusOK, bfsResponse{ // pake bfsResponse aja biar simple
			Target:       target,
			VisitedCount: res.VisitedCount,
			Trees:        res.Trees,
		})
	})

	r.GET("/api/bdr", func(c *gin.Context) {
		target := c.Query("target")
		if target == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "query param 'target' is required"})
			return
		}

		maxRecipe, _ := strconv.Atoi(c.DefaultQuery("max_recipe", "0"))

		// panggil fungsi BDR dari backend
		res := backend.BDR(gallery, target, maxRecipe)

		// kembalikan dengan format yang sama seperti BFS/DFS
		c.JSON(http.StatusOK, bfsResponse{
			Target:       target,
			VisitedCount: res.VisitedCount,
			Trees:        res.Trees,
		})
	})

	r.GET("/api/elements", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"elements": gallery.GetAllNames(), // misal fungsi helper di backend
		})
	})

	r.GET("/api/bfs/stream", func(c *gin.Context) {
		target := c.Query("target")
		if target == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "'target' is required"})
			return
		}
		maxRecipe, _ := strconv.Atoi(c.DefaultQuery("max_recipe", "0"))

		w := c.Writer
		flusher, ok := w.(http.Flusher)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "streaming not supported"})
			return
		}
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		results := make(chan *backend.RecipeNode)
		go func() {
			backend.BFSStream(gallery, target, maxRecipe, results)
			close(results)
		}()

		enc := json.NewEncoder(w)
		for node := range results {
			fmt.Fprint(w, "data: ")
			if err := enc.Encode(node); err != nil {
				log.Println("encode:", err)
				break
			}
			fmt.Fprint(w, "\n")
			flusher.Flush()
		}
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
