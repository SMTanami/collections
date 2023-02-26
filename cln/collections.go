package cln

// A generic Collection interface for common data structures. The size of the interface is subject to change as interface composition changes/improves overtime.
type Collection[T comparable] interface {
	Add(vals ...T)
	Take() (T, bool)
	Contains(val T) bool
	Remove(val T)
	Filter(filter func(v T) bool)
	Clear()
	Size() int
	IsEmpty() bool
	Iter() chan T
	String() string
}
