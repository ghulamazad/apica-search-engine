package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ghulamazad/GFileMux"
	"github.com/ghulamazad/GFileMux/storage"
	"github.com/ghulamazad/apica-search-engine/api"
	"github.com/ghulamazad/apica-search-engine/indexer"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func getGFileMuxMiddleware() (*GFileMux.GFileMux, *storage.DiskStorage, error) {
	// Initialize disk storage
	disk, err := storage.NewDiskStorage("data")
	if err != nil {
		log.Fatalf("Error initializing disk storage: %v", err)
		return nil, nil, err
	}

	// Create a file handler with desired configurations
	uploader, err := GFileMux.New(
		GFileMux.WithMaxFileSize(10<<20), // Limit file size to 10MB
		GFileMux.WithFileValidatorFunc(
			GFileMux.ChainValidators(GFileMux.ValidateMimeType("parquet", "application/octet-stream")),
		),
		GFileMux.WithFileNameGeneratorFunc(func(originalFileName string) string {
			// Generate a new unique file name using UUID and original file extension
			parts := strings.Split(originalFileName, ".")
			// handle without extension file
			if len(parts) == 1 {
				return fmt.Sprintf("%s_%s", uuid.NewString(), originalFileName)
			}
			// handle with extension file
			ext := parts[len(parts)-1]
			return fmt.Sprintf("%s.%s", uuid.NewString(), ext)
		}),
		GFileMux.WithStorage(disk), // Use disk storage
	)

	if err != nil {
		log.Fatalf("Error initializing file handler: %v", err)
		return nil, nil, err
	}

	return uploader, disk, nil
}

func main() {
	idx := indexer.NewInvertedIndex()

	r := mux.NewRouter()

	// File upload middleware
	gFileMux, storageDisk, err := getGFileMuxMiddleware()
	if err != nil {
		log.Fatalf("Error initializing GFileMux: %v", err)
	}

	h := &api.SearchHandler{
		Indexer: idx,
		Disk:    storageDisk,
	}

	// load initial files
	h.LoadFiles([]string{"data/File 1", "data/File 2"})

	// Search handler
	r.HandleFunc("/search", h.Search).Methods("GET")

	// Upload handler
	r.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		gFileMux.Upload("bucket_name", "files")(http.HandlerFunc(h.UploadHandler)).ServeHTTP(w, r)
	}).Methods("POST")

	// Wrap router with CORS middleware
	handler := cors.AllowAll().Handler(r)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
