package elastic

import (
	"context"
	"fmt"
	pkgElastic "github.com/olivere/elastic/v7"
	"log"
)

// ElasticContract is an interface that defines methods for interacting with an ElasticSearch backend.
type ElasticContract interface {
	Store(indexName string, data interface{}, ctx context.Context) error
	Get(ctx context.Context, indexName, queryName, queryContent string) (*pkgElastic.SearchResult, error)
}

// elastic is a struct that represents an ElasticSearch client.
// It holds a pointer to a pkgElastic.Client for communication with the ElasticSearch backend.
type elastic struct {
	client *pkgElastic.Client
}

// NewElastic is a function that creates and returns a new instance of the ElasticContract interface.
// It takes a pointer to a pkgElastic.Client as a parameter and initializes an 'elastic' struct with it.
// The 'elastic' struct is returned as an implementation of the ElasticContract interface.
func NewElastic(client *pkgElastic.Client) ElasticContract {
	return &elastic{
		client: client,
	}
}

// Store stores the provided data in the specified index.
// It takes an indexName string, data interface{}, and a context.Context as parameters.
// Returns an error if the storing process encounters any issues.
func (e *elastic) Store(indexName string, data interface{}, ctx context.Context) error {
	_, err := e.client.Index().Index(indexName).BodyJson(data).Do(ctx)
	if err != nil {
		log.Fatalf("Error indexing document: %v", err)
		return err
	}

	fmt.Println("Document indexed successfully")

	return nil
}

// Get retrieves data from the specified index based on the provided query.
// It takes a context.Context, indexName string, queryName string, and queryContent string as parameters.
// Returns a pointer to pkgElastic.SearchResult and an error if the retrieval process encounters any issues.
func (e *elastic) Get(ctx context.Context, indexName, queryName, queryContent string) (*pkgElastic.SearchResult, error) {
	// Search for documents
	searchResult, err := e.client.Search().Index(indexName).Query(pkgElastic.NewMatchQuery(queryName, queryContent)).Do(ctx)
	if err != nil {
		log.Fatalf("Error searching documents: %v", err)
		return nil, err
	}

	fmt.Printf("Found %d documents\n", searchResult.TotalHits())
	return searchResult, nil
}
