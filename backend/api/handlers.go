package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ghulamazad/GFileMux"
	"github.com/ghulamazad/GFileMux/storage"
	"github.com/ghulamazad/apica-search-engine/indexer"
)

type SearchHandler struct {
	Indexer *indexer.InvertedIndex
	Disk    *storage.DiskStorage
}

func (h *SearchHandler) LoadFiles(files []string) {
	for _, file := range files {
		records, err := indexer.ParseParquet(file)
		if err != nil {
			log.Fatalf("Error parsing file %s: %v", file, err)
		}

		h.Indexer.AddRecords(records)
	}
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

func (h *SearchHandler) UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the uploaded files from the request context
	files, err := GFileMux.GetUploadedFilesFromContext(r)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, fmt.Sprintf("Failed to get uploaded files: %v", err))))
		return
	}

	var uploadedFiles []string
	// Process each uploaded file and print details
	for _, file := range files {
		// Print the file path in disk storage
		filePath, err := h.Disk.Path(context.Background(), GFileMux.PathOptions{
			Key:    file[0].StorageKey,
			Bucket: file[0].FolderDestination,
		})
		if err != nil {
			log.Printf("Error retrieving file path for %s: %v", file[0].StorageKey, err)
			continue // Skip to the next file if there's an error
		}

		uploadedFiles = append(uploadedFiles, filePath)
	}

	// load the uploaded files
	h.LoadFiles(uploadedFiles)

	// Respond with a success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Files uploaded successfully"}`))

}
