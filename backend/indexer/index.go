package indexer

import (
	"sync"

	"github.com/ghulamazad/apica-search-engine/models"
	"github.com/ghulamazad/apica-search-engine/utils"
)

type InvertedIndex struct {
	mu      sync.RWMutex
	index   map[string][]*models.Record
	records []*models.Record
}

func NewInvertedIndex() *InvertedIndex {
	return &InvertedIndex{
		index: make(map[string][]*models.Record),
	}
}

func (ii *InvertedIndex) AddRecords(records []*models.Record) {
	ii.mu.Lock()
	defer ii.mu.Unlock()

	for _, rec := range records {
		ii.records = append(ii.records, rec)
		content := rec.Message + " " + rec.MessageRaw + " " + rec.Tag + " " + rec.StructuredData
		tokens := utils.Tokenize(content)
		for _, token := range tokens {
			ii.index[token] = append(ii.index[token], rec)
		}
	}
}

func (ii *InvertedIndex) Search(query string) []*models.Record {
	ii.mu.RLock()
	defer ii.mu.RUnlock()

	results := make(map[*models.Record]struct{})
	for _, token := range utils.Tokenize(query) {
		for _, record := range ii.index[token] {
			results[record] = struct{}{}
		}
	}

	var out []*models.Record
	for rec := range results {
		out = append(out, rec)
	}

	return out
}
