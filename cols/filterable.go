package cols

// Filter is the interface that wraps the basic Filter and Remove methods.
type Filterable[T comparable] interface {
	Remove(val T)
	Filter(filter func(v T) bool)
}
