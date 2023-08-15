package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go_elastic/config"
	e "go_elastic/elastic"
	"go_elastic/entity"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

func main() {
	var c config.Config
	cfg, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatal(fmt.Sprintf("error reading file yaml err : %s", err))
	}

	if err := yaml.Unmarshal(cfg, &c); err != nil {
		panic(err)
	}

	client := config.ElasticConnection(c)
	ctx := context.Background()
	es := e.NewElastic(client)

	// Index a post
	indexName := "social-media-post"
	post := entity.Post{
		Text:    "Excited about my new project!",
		UserID:  "user123",
		Created: time.Now(),
	}

	// Attempt to store the 'post' data in the specified 'indexName' using the 'Store' method of the 'es' ElasticContract instance.
	// If an error occurs during the indexing process, log a fatal error message with the details and terminate the program.
	err = es.Store(indexName, post, ctx)
	if err != nil {
		log.Fatalf("Error indexing document: %v", err)
	}
	fmt.Println("Document indexed successfully")

	// Attempt to retrieve documents from the specified 'indexName' using the 'Get' method of the 'es' ElasticContract instance.
	// If an error occurs during the search process, log a fatal error message with the details and terminate the program.
	sr, err := es.Get(ctx, indexName, "text", "great")
	if err != nil {
		log.Fatalf("Error searching documents: %v", err)
	}

	fmt.Printf("Found %d documents\n", sr.TotalHits())

	var posts []entity.Post

	// Print document for variable sr
	for _, hit := range sr.Hits.Hits {
		fmt.Printf("Document ID: %s, Score: %f\n", hit.Id, hit.Score)

		var post entity.Post
		err := json.Unmarshal(hit.Source, &post)
		if err != nil {
			log.Printf("Error unmarshaling document: %v", err)
			continue
		}

		posts = append(posts, post)
	}
	//Print all data posts
	for _, e := range posts {
		fmt.Println("Text", e.Text)
		fmt.Println("User", e.UserID)
		fmt.Println("Created", e.Created)
	}

	// Close the client when done
	defer client.Stop()
}
