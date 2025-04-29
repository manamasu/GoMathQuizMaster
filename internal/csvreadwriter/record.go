package csvreadwriter

// Record is a generic container for a row of data, which can be converted to/from CSV
type Record[T any] struct {
	Record []T
}

//Constructor: NewRecord creates a new Record[T] from a slice of T
func NewRecord[T any](data []T) Record[T] {
	return Record[T]{Record: data}
}
