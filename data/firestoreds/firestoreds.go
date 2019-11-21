package firestoreds

import (
	"context"
	"fmt"

	"todobackend-gcr/data/firestoreds/todosds"

	"cloud.google.com/go/firestore"
)

// Store contains all the DataStores for the different resources.
type Store struct {
	Todos *todosds.DataStore
}

// NewDataStore creates a new DataStore.
func NewDataStore(ctx context.Context, client *firestore.Client) (*Store, error) {
	var ds Store
	// Create the datastore for each individual resource.
	todosds, err := todosds.NewDataStore(ctx, client)
	if err != nil {
		return &ds, fmt.Errorf("error creating new todos data store: %s", err)
	}

	// Add all the resource datastores to the overall datastore. In this example,
	// we only have the todos datastore.
	ds.Todos = todosds

	return &ds, nil
}
