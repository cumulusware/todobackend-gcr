package todos

// DataStore provides the interface required to retrieve and save todos.
type DataStore interface {
	Create(*Todo) (string, error)
	GetAll() ([]Todo, error)
	DeleteAll() error
}
