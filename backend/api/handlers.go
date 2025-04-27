package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ghulamazad/apica-search-engine/indexer"
)

type SearchHandler struct {
	Indexer *indexer.InvertedIndex
}

func (h *SearchHandler) Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	start := time.Now()
	results := h.Indexer.Search(query)
	duration := time.Since(start)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"results":  results,
		"count":    len(results),
		"duration": duration.String(),
	})
}
