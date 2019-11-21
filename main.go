package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"todobackend-gcr/data/firestoreds"

	"cloud.google.com/go/firestore"
)

const (
	projectID = "todobackendgcr"
)

func main() {
	// Get port or default to 8080.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Get a Firestore client.
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create firestore client: %v", err)
	}

	// Close client when done.
	defer client.Close()

	// Create the data store.
	ds, err := firestoreds.NewDataStore(ctx, client)
	if err != nil {
		log.Fatalf("error creating datastore: %s", err)
	}

	// Create routes with injected data store and CORS. Then start server.
	r := createRoutes(ds)
	c := setupCors()
	log.Fatal(http.ListenAndServe(":"+port, c.Handler(r)))
}
