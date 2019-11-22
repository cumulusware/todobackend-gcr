package todosds

import (
	"context"
	"log"

	"todobackend-gcr/resources/todos"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// DataStore implements the DataStore interface for todos.
type DataStore struct {
	ctx        context.Context
	Collection *firestore.CollectionRef
	Client     *firestore.Client
}

// NewDataStore creates a new DataStore.
func NewDataStore(ctx context.Context, c *firestore.Client) (*DataStore, error) {
	var ds DataStore
	dbName := "todos"
	ref := c.Collection(dbName)
	ds = DataStore{ctx, ref, c}
	return &ds, nil
}

// GetAll returns all todos found in the DataStore.
func (ds *DataStore) GetAll(baseURL string) ([]todos.Todo, error) {
	var theTodos []todos.Todo

	iter := ds.Collection.Documents(ds.ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		var todo todos.Todo
		err = doc.DataTo(&todo)
		if err != nil {
			return theTodos, err
		}
		todo.URL = baseURL + doc.Ref.ID
		theTodos = append(theTodos, todo)
	}

	return theTodos, nil
}

// GetByID returns one todo found in the DataStore.
func (ds *DataStore) GetByID(id, url string) (todos.Todo, error) {
	var todo todos.Todo
	docsnap, err := ds.Collection.Doc(id).Get(ds.ctx)
	if err != nil {
		return todo, err
	}
	err = docsnap.DataTo(&todo)
	todo.URL = url
	return todo, nil
}

// Create stores a new todo in the DataStore.
func (ds *DataStore) Create(todo *todos.Todo) (string, error) {
	docRef, _, err := ds.Collection.Add(ds.ctx, todo)
	return docRef.ID, err
}

// DeleteAll deletes all todos in the DataStore.
func (ds *DataStore) DeleteAll() error {

	// Get all docs.
	iter := ds.Collection.Documents(ds.ctx)
	numToDelete := 0
	batch := ds.Client.Batch()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		batch.Delete(doc.Ref)
		numToDelete++
	}
	if numToDelete == 0 {
		// Nothing to delete, let's leave.
		return nil
	}
	_, err := batch.Commit(ds.ctx)
	return err
}

// DeleteByID delets one todo found in the DataStore.
func (ds *DataStore) DeleteByID(id string) error {
	_, err := ds.Collection.Doc(id).Delete(ds.ctx)
	return err
}

// UpdateByID updates one todo found in the DataStore.
func (ds *DataStore) UpdateByID(id string, todo *todos.Todo) error {
	_, err := ds.Collection.Doc(id).Set(ds.ctx, todo)
	return err
}
