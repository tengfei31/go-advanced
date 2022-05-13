package db

import (
	// "github.com/elastic/go-elasticsearch/v7"
	"github.com/olivere/elastic/v7"
)

var EsClient *elastic.Client

func initEsClient() error {
	var err error
	EsClient, err = elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		return err
	}
	return nil
}

func init() {
	if err := initEsClient(); err != nil {
		panic(err)
	}
}
