package main

import (
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

		path := backend.DFS(gallery, target, backend.DFSOptions{}) // misal fungsi DFS mengembalikan []string
		c.JSON(http.StatusOK, gin.H{
			"target": target,
			"path":   path,
		})
	})

	r.GET("/api/elements", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"elements": gallery.GetAllNames(), // misal fungsi helper di backend
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
