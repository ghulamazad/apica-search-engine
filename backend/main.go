package main

import (
	"context"
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

func loadFiles(idx *indexer.InvertedIndex, files []string) {
	for _, file := range files {
		records, err := indexer.ParseParquet(file)
		if err != nil {
			log.Fatalf("Error parsing file %s: %v", file, err)
		}

		idx.AddRecords(records)
	}
}

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

	files := []string{
		"data/File 1",
		"data/File 2",
		"data/File 3",
		"data/File 4",
		"data/File 5",
		"data/File 6",
		"data/File 7",
		"data/File 8",
		"data/File 9",
		"data/File 10",
		"data/File 11",
		"data/File 14",
		"data/File 15",
		"data/File 16",
	}

	loadFiles(idx, files)

	r := mux.NewRouter()

	// File upload middleware
	gFileMux, storageDisk, err := getGFileMuxMiddleware()
	if err != nil {
		log.Fatalf("Error initializing GFileMux: %v", err)
	}

	// Search handler
	h := &api.SearchHandler{
		Indexer: idx,
	}
	r.HandleFunc("/search", h.Search).Methods("GET")

	// Upload handler
	r.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		gFileMux.Upload("bucket_name", "files")(http.HandlerFunc(uploadHandler(storageDisk, idx))).ServeHTTP(w, r)
	}).Methods("POST")

	// Wrap router with CORS middleware
	handler := cors.AllowAll().Handler(r)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, message)))
}

func uploadHandler(disk *storage.DiskStorage, idx *indexer.InvertedIndex) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the uploaded files from the request context
		files, err := GFileMux.GetUploadedFilesFromContext(r)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get uploaded files: %v", err))
			return
		}

		var uploadedFiles []string
		// Process each uploaded file and print details
		for _, file := range files {
			// Print the file path in disk storage
			filePath, err := disk.Path(context.Background(), GFileMux.PathOptions{
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
		loadFiles(idx, uploadedFiles)

		// Respond with a success message
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Files uploaded successfully"}`))
	}
}
