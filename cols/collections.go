package cols

type Collection[T comparable] interface {
	Add(vals ...T)
	Take() (T, bool)
	Contains(val T) bool
	Clear()
	Size() int
	IsEmpty() bool
}
