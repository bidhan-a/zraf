package store

// Store defines the storage mechanism for the graph database.
type Store interface {
}

// Transaction represents a database transaction.
type Transaction interface {
	Add()
	Commit()
	Rollback()
}
