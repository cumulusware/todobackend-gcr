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

// Create stores a new todo in the DataStore.
func (ds *DataStore) Create(todo *todos.Todo) (string, error) {
	docRef, _, err := ds.Collection.Add(ds.ctx, todo)
	return docRef.ID, err
}

// GetAll returns all todos found in the DataStore.
func (ds *DataStore) GetAll() ([]todos.Todo, error) {
	var todos []todos.Todo

	iter := ds.Collection.Documents(ds.ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		todo := convertDocToTodo(doc.Data())
		todos = append(todos, todo)
	}

	return todos, nil
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

func convertDocToTodo(doc map[string]interface{}) todos.Todo {
	var todo todos.Todo
	title, ok := doc["title"].(string)
	if !ok {
		todo.Title = ""
	} else {
		todo.Title = title
	}
	return todo
}
