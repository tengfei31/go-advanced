package distribute

import (
	"context"
	"go-advanced/db"
	"strings"

	"github.com/olivere/elastic/v7"
)

func insertDocument(dbName string, table string, obj map[string]interface{}) error {
	id := obj["id"].(string)

	var indexName, typeName string
	indexName = strings.Join([]string{dbName, table}, "_")
	typeName = table

	_, err := db.EsClient.Index().Index(indexName).Type(typeName).Id(id).Do(context.Background())
	if err != nil {
		return err
	}
	return nil

}

func query(indexName string, typeName string) (*elastic.SearchResult, error) {
	q := elastic.NewBoolQuery().Must(elastic.NewMatchPhraseQuery("id", 1), elastic.NewBoolQuery().Must(elastic.NewMatchPhraseQuery("male", "m")))

	q = q.Should(elastic.NewMatchPhraseQuery("name", "alex"), elastic.NewMatchPhraseQuery("name", "xargin"))

	searchService := db.EsClient.Search(indexName).Type(typeName)
	res, err := searchService.Query(q).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return res, nil
}

func deleteDocument(indexName string, typeName string, obj map[string]interface{}) error {
	id := obj["id"].(string)

	_, err := db.EsClient.Delete().Index(indexName).Type(typeName).Id(id).Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
