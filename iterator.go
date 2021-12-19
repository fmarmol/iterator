package iterator

type Iterator[T any] interface {
	Iter() <-chan T
}
