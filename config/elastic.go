package config

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
)

func ElasticConnection(cfg Config) *elastic.Client {
	//Create an Elasticsearch client
	client, err := elastic.NewClient(
		elastic.SetURL(cfg.ElasticURL), // Elasticsearch server URL
		elastic.SetSniff(false),        // Disable automatic node discovery
	)

	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %v", err)
	}

	//Ping the Elasticsearch server to check its availability
	info, code, err := client.Ping(cfg.ElasticURL).Do(context.Background())
	if err != nil {
		log.Fatalf("Error pinging Elasticsearch: %v", err)
	}

	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	return client
}
