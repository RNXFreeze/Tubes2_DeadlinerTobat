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
/* Deskripsi  : F00A - Main Program API (Connection Frontend & Backend)          */
/* PIC F00A   : K01 - 13523050 - Mayla Yaffa Ludmilla                            */

package main

import (
	"fmt"
	"log"
	"strconv"
	"net/http"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"Tubes2_DeadlinerTobat/src/backend"
)

type BFSResponse struct {
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
		res := backend.BFS(gallery , target , backend.AlgorithmOption{MaxRecipes : maxRecipe , LiveChan : nil});

		c.JSON(http.StatusOK, BFSResponse{
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

		res := backend.DFS(gallery , target , backend.AlgorithmOption{MaxRecipes : maxRecipe , LiveChan : nil});

		c.JSON(http.StatusOK, BFSResponse{ // pake BFSResponse aja biar simple
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
		res := backend.BDR(gallery , target , backend.AlgorithmOption{MaxRecipes : maxRecipe , LiveChan : nil});

		// kembalikan dengan format yang sama seperti BFS/DFS
		c.JSON(http.StatusOK, BFSResponse{
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
			option := backend.AlgorithmOption{
				MaxRecipes: maxRecipe,
				LiveChan:  results,
			}
			fmt.Println("MaxRecipes =", maxRecipe)
			backend.BFSStream(gallery, target, option.MaxRecipes, option.LiveChan)
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

	r.GET("/api/dfs/stream", func(c *gin.Context) {
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
			option := backend.AlgorithmOption{
				MaxRecipes: maxRecipe,
				LiveChan:  results,
			}
			backend.DFSStream(gallery, target, option)
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

	r.GET("/api/bdr/stream", func(c *gin.Context) {
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
			option := backend.AlgorithmOption{
				MaxRecipes: maxRecipe,
				LiveChan:  results,
			}
			backend.BDRStream(gallery, target, option)
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
