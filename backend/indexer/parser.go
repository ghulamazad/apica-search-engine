package indexer

import (
	"log"

	"github.com/ghulamazad/apica-search-engine/models"
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/reader"
)

func ParseParquet(filepath string) ([]*models.Record, error) {
	fr, err := local.NewLocalFileReader(filepath)
	if err != nil {
		log.Println("Can't open file")
		return nil, err
	}
	defer fr.Close()

	pr, err := reader.NewParquetReader(fr, new(models.Record), 4)
	if err != nil {
		log.Println("Can't create parquet reader", err)
		return nil, err
	}
	defer pr.ReadStop()

	num := int(pr.GetNumRows())
	records := make([]*models.Record, num)

	if err := pr.Read(&records); err != nil {
		return nil, err
	}

	log.Printf("Loaded %d records from %s\n", len(records), filepath)
	return records, nil
}
