package cols

type Iterable[T comparable] interface {
	Iter() chan T
}
