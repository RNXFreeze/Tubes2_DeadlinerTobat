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
/* Deskripsi  : F00 - Main Program API (Connection Frontend & Backend)           */
/* PIC F00    : K01 - 13523050 - Mayla Yaffa Ludmilla                            */

package main

import (
	"log"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"Tubes2_DeadlinerTobat/src/backend"
)

type AlgorithmResponse struct {
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

		backend.EnableMultithreading();
		maxRecipe, _ := strconv.Atoi(c.DefaultQuery("max_recipe", "0"))
		res := backend.BFS(gallery, target, maxRecipe)

		c.JSON(http.StatusOK, AlgorithmResponse{
			Target:       target,
			VisitedCount: res.VisitedCount,
			Trees:        res.Trees, // sudah JSON-marshal-able karena ada tag di RecipeNode
		})
		backend.DisableMultithreading();
	})
	r.GET("/api/dfs", func(c *gin.Context) {
		target := c.Query("target")
		if target == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "query param 'target' is required"})
			return
		}

		backend.EnableMultithreading();
		maxRecipe, _ := strconv.Atoi(c.DefaultQuery("max_recipe", "0"))

		res := backend.DFS(gallery, target, maxRecipe)
		backend.DisableMultithreading();

		c.JSON(http.StatusOK, AlgorithmResponse{ // pake AlgorithmResponse aja biar simple
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

		backend.EnableMultithreading();
		maxRecipe, _ := strconv.Atoi(c.DefaultQuery("max_recipe", "0"))

		res := backend.BDR(gallery, target, maxRecipe)
		backend.DisableMultithreading();

		c.JSON(http.StatusOK, AlgorithmResponse{
			Target:       target,
			VisitedCount: res.VisitedCount,
			Trees:        res.Trees,
		})
	})

	r.GET("/api/elements", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"elements": gallery.GetAllNames(),
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}